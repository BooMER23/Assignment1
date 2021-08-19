package HTTP

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BooMER23/golang.git/DB"
)

type Final_list struct{
	Item string 		`json:"Item"`
	Quantity float64   `json:"Quantity"`
}

var List Final_list
var Lists []Final_list 

func fetchingremainingitems(){					//to fetch the data from tabel-"remaining items"
	PSQLserver := DB.ConnectServer()

	db, err := sql.Open("postgres", PSQLserver)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected successfully!")
	}
	defer db.Close()
	Rows, err := db.Query("SELECT * FROM remainingitems")
	if err != nil {
		panic(err)
	}
	for Rows.Next() {
		Rows.Scan(&List.Item,&List.Quantity)
		Lists = append(Lists,List)
	}
}

func RemainingItems(w http.ResponseWriter, r *http.Request){
	fetchingremainingitems()
	json.NewEncoder(w).Encode(Lists)
}