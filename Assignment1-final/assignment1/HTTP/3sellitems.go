package HTTP

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BooMER23/golang.git/DB"
)

var Buy map[string]interface{}	//this will store the request body

type BuyJson struct {
	Name     string  `json:"Name"`
	Item     string  `json:"Item"`
	Quantity float64 `json:"Quantity"`
}

var MulBuy []BuyJson



func SellItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	conv := Buy
	json.NewDecoder(r.Body).Decode(&conv)
	item := make(map[string]interface{})
	quantity := make(map[string]interface{})
	var buy BuyJson
	for i := range conv {
		item[i] = conv[i].(map[string]interface{})["Item"]
		quantity[i] = conv[i].(map[string]interface{})["Quantity"]
		x := i
		y := item[i].(string)
		z := quantity[i].(float64)
		buy.Name = x
		buy.Item = y
		buy.Quantity = z
		MulBuy = append(MulBuy, buy)
	}
	infoOfCustomer()	
}


func infoOfCustomer(){							//inserting information about customer into table-"buy"
	PSQLserver := DB.ConnectServer()

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