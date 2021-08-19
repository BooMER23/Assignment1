package HTTP

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/BooMER23/golang.git/DB"
)

type remitem struct {
	Item     string  `json:"Item"`
	Quantity float64 `json:"Quantity"`
}

var Final remitem

var UpdatedItems []remitem

func Itemsremaining(w http.ResponseWriter, r *http.Request) {
	for index := range MulBuy {
		for i := range products {
			if MulBuy[index].Item == products[i].Item {
				Final.Item = MulBuy[index].Item
				Final.Quantity = products[i].Quantity - MulBuy[index].Quantity
				products[i].Quantity = Final.Quantity
			}
		}
		UpdatedItems = append(UpdatedItems, Final)
	}
	insertintoremainingitems()
}

func insertintoremainingitems(){										//insert into the table remaining items
	PSQLserver := DB.ConnectServer()

	db, err := sql.Open("postgres", PSQLserver)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected successfully!")
	}
	defer db.Close()
	for index := range UpdatedItems {
		sqlStatement := `INSERT INTO remainingitems (item,remaining_quantity) VALUES ($1,$2)`
		_,err := db.Exec(sqlStatement,UpdatedItems[index].Item, UpdatedItems[index].Quantity)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("\nRow inserted successfully!")
		}
	}
}