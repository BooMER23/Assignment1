package DB

import (
	"database/sql"
	"fmt"
)

type CustomerDetails struct{
	Name string `json:"Name"`
	Item string `json:"Item"`
	Quantity float64 `json:"Quantity"`
}

var MulCustomerDeatils []CustomerDetails

func FetchsoldItems() []CustomerDetails{			//this will fetch the details of the customer
	PSQLserver := ConnectServer()

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
		var customer CustomerDetails
		Rows.Scan(&customer.Name,&customer.Item, &customer.Quantity)
		MulCustomerDeatils = append(MulCustomerDeatils, customer)
	}
	return MulCustomerDeatils
}