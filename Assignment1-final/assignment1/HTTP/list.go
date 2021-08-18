package HTTP

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/BooMER23/golang.git/DB"
)

type Finall struct {
	Name     string  `json:"Name"`
	Item     string  `json:"Item"`
	Quantity float64 `json:"Quantity"`
}

var Final Finall

var UpdatedItems []Finall

func FinalList(w http.ResponseWriter, r *http.Request) {
	PSQLserver := DB.ConnectServer()

	db, err := sql.Open("postgres", PSQLserver)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected successfully!")
	}
	defer db.Close()
	for index := range UpdatedItems {
		sqlStatement := `INSERT INTO final (name,item,quantity) VALUES ($1,$2,$3)`
		_,err := db.Exec(sqlStatement,UpdatedItems[index].Name,UpdatedItems[index].Item,UpdatedItems[index].Quantity)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("\nRow inserted successfully!")
		}
	}
}