package DB

import (
	"database/sql"
	"fmt"
	"time"
)

type Row struct {
	ID          int       `json:"ID"`
	Item        string    `json:"Item"`
	Price       float64   `json:"Price"`
	Quantity    float64   `json:"Quantity"`
	Create_Time time.Time `json:"Create_Time"`
}
var Erows []Row

func FectchingallItems() []Row {				//this func will fetch allItems from the database
	PSQLserver := ConnectServer()

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
		Erows = append(Erows, db_row)
	}
	return Erows
}