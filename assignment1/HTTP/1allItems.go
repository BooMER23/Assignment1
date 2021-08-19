package HTTP

import (
	"net/http"
	"encoding/json"
	"time"
	"github.com/BooMER23/golang.git/DB"
)

var Items map[string]interface{}				//The resquest will store in the Items interface

func AllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	conv := Items
	json.NewDecoder(r.Body).Decode(&conv)
	values1 := make(map[string]interface{})
	values2 := make(map[string]interface{})
	//var product Product
	id := 0
	for i := range conv {
		id++
		values1[i] = (conv[i]).(map[string]interface{})["Price"]
		values2[i] = (conv[i]).(map[string]interface{})["Quantity"]
		x := values1[i].(float64)
		y := values2[i].(float64)
		z := x / y
		values1[i] = z
		DB.Eproduct.ID = id						//the values directly uploaded to the database insert all Items
		DB.Eproduct.Item = i
		DB.Eproduct.PriceFortheQuantity = x
		DB.Eproduct.Quantity = y
		DB.Eproduct.Price = z
		DB.Eproduct.Time = time.Now()
		DB.Eproducts = append(DB.Eproducts, DB.Eproduct)
	}
	DB.InsertintoallItems()							
}