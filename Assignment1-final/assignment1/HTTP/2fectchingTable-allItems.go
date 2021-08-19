package HTTP

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/BooMER23/golang.git/DB"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)



type Row struct {
	ID          int       `json:"ID"`
	Item        string    `json:"Item"`
	Price       float64   `json:"Price"`
	Quantity    float64   `json:"Quantity"`
	Create_Time time.Time `json:"Create_Time"`
}
var rows []Row

func TofetchfromallItems() {					//fetching table allItems
	PSQLserver := DB.ConnectServer()

	db, err := sql.Open("postgres", PSQLserver)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected successfully!")
	}
	defer db.Close()
	Rows, err := db.Query("SELECT * FROM allItems")
	if err != nil {
		panic(err)
	}

	for Rows.Next() {
		var db_row Row
		Rows.Scan(&db_row.ID, &db_row.Item, &db_row.Price, &db_row.Quantity, &db_row.Create_Time)
		rows = append(rows, db_row)
	}
}

func FetchingallItems(w http.ResponseWriter, r *http.Request) {
	TofetchfromallItems()
	json.NewEncoder(w).Encode(rows)
}