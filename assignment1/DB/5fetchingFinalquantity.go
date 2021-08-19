package DB

import (
	"database/sql"
	"fmt"
)

type remitem struct {
	Item     string  `json:"Item"`
	Quantity float64 `json:"Quantity"`
}

var Final remitem

var UpdatedItems []remitem

func Insertintoremainingitems(){						//insert into the table remaining items
	
	for index := range MulBuy {
		for i := range Eproducts {
			if MulBuy[index].Item == Eproducts[i].Item {
				Final.Item = MulBuy[index].Item
				Final.Quantity = Eproducts[i].Quantity - MulBuy[index].Quantity
				Eproducts[i].Quantity = Final.Quantity
			}
		}
		UpdatedItems = append(UpdatedItems, Final)
	}
	
	PSQLserver := ConnectServer()

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