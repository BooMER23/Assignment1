package DB

import (
	"database/sql"
	"fmt"
)

type BuyJson struct {
	Name     string  `json:"Name"`
	Item     string  `json:"Item"`
	Quantity float64 `json:"Quantity"`
}
var Ebuy BuyJson

var MulBuy []BuyJson

func InfoOfCustomer(){							//inserting information about customer into table-"buy"
	PSQLserver := ConnectServer()

	db, err := sql.Open("postgres", PSQLserver)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected successfully!")
	}
	defer db.Close()
	for i := range MulBuy {
		sqlStatement := `INSERT INTO customerdetails (name,item,quantity) VALUES ($1,$2,$3)`
		_, err = db.Exec(sqlStatement, MulBuy[i].Name, MulBuy[i].Item, MulBuy[i].Quantity)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("\nRow inserted successfully!")
		}
	}
}