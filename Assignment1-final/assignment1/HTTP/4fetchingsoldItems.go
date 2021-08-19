package HTTP

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BooMER23/golang.git/DB"
)

type CustomerDeatils struct{
	Name string `json:"Name"`
	Item string `json:"Item"`
	Quantity float64 `json:"Quantity"`
}

var MulCustomerDeatils []CustomerDeatils

func fetchsoldItems(){
	PSQLserver := DB.ConnectServer()

	db, err := sql.Open("postgres", PSQLserver)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected successfully!")
	}
	defer db.Close()
	Rows, err := db.Query("SELECT * FROM customerdetails")
	if err != nil {
		panic(err)
	}

	for Rows.Next() {
		var customer CustomerDeatils
		Rows.Scan(&customer.Name,&customer.Item, &customer.Quantity)
		MulCustomerDeatils = append(MulCustomerDeatils, customer)
	}
}

func SoldItems(w http.ResponseWriter, r *http.Request){			//this will fetch the details of the customers
	fetchsoldItems()
	json.NewEncoder(w).Encode(MulCustomerDeatils)
}



