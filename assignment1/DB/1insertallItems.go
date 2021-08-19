package DB

import(
	"fmt"
	"database/sql"
	"time"
)

type Product struct {
	ID                  int       `json:"ID"`
	Item                string    `json:"Item"`
	PriceFortheQuantity float64   `json:"PriceFortheQuantity"`
	Quantity            float64   `json:"Quantity"`
	Price               float64   `json:"Price"`
	Time                time.Time `json:"Time"`
}
var Eproduct Product			//Eproduct will export the struct product to allItems in HTTP
var Eproducts []Product

func InsertintoallItems(){
	PSQLserver := ConnectServer()

	db, err := sql.Open("postgres", PSQLserver)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected successfully!")
	}
	defer db.Close()
	for index := range Eproducts {
		sqlStatement := `INSERT INTO allItems (item,price,quantity,create_time)
		VALUES ($1,$2,$3,$4)`
		_, err = db.Exec(sqlStatement, string(Eproducts[index].Item), float64(Eproducts[index].Price), float64(Eproducts[index].Quantity), (Eproducts[index].Time))
		if err != nil {
			panic(err)
		} else {
			fmt.Println("\nRow inserted successfully!")
		}
	}
}