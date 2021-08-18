package HTTP

import (
	"net/http"
	"fmt"
	"encoding/json"
	"time"
	"database/sql"
	"github.com/BooMER23/golang.git/DB"
)

var Items map[string]interface{}
type Product struct {
	ID                  int       `json:"ID"`
	Name                string    `json:"Name"`
	PriceFortheQuantity float64   `json:"PriceFortheQuantity"`
	Quantity            float64   `json:"Quantity"`
	Price               float64   `json:"Price"`
	Time                time.Time `json:"Time"`
}
var products []Product

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	conv := Items
	json.NewDecoder(r.Body).Decode(&conv)
	values1 := make(map[string]interface{})
	values2 := make(map[string]interface{})
	var product Product
	id := 0
	for i := range conv {
		id++
		values1[i] = (conv[i]).(map[string]interface{})["Price"]
		values2[i] = (conv[i]).(map[string]interface{})["Quantity"]
		x := values1[i].(float64)
		y := values2[i].(float64)
		z := x / y
		values1[i] = z
		product.ID = id
		product.Name = i
		product.PriceFortheQuantity = x
		product.Quantity = y
		product.Price = z
		product.Time = time.Now()
		products = append(products, product)
	}
	PSQLserver := DB.ConnectServer()

	db, err := sql.Open("postgres", PSQLserver)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected successfully!")
	}
	defer db.Close()
	for index := range products {
		sqlStatement := `INSERT INTO infor (item,price,quantity,create_time)
		VALUES ($1,$2,$3,$4)`
		_, err = db.Exec(sqlStatement, string(products[index].Name), float64(products[index].Price), float64(products[index].Quantity), (products[index].Time))
		if err != nil {
			panic(err)
		} else {
			fmt.Println("\nRow inserted successfully!")
		}
	}
}