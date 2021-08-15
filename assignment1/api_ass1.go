package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"database/sql"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/postgres"
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

type Row struct {
	ID int `json:"ID"`
	Item string `json:"Item"`
	Price float64 `json:"Price"`
	Quantity float64 `json:"Quantity"`
	Create_Time time.Time `json:"Create_Time"`
}


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "abc123"
	dbname   = "provisions"
)

var products []Product

func main() {
	router := mux.NewRouter()

	for index := range products {
		fmt.Println(products[index])
	}

	router.HandleFunc("/", HomePage).Methods("POST")
	router.HandleFunc("/rows", rows).Methods("GET")
	http.ListenAndServe(":8080", router)
}

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
		product.Name = i //interface to struct
		product.PriceFortheQuantity = x
		product.Quantity = y
		product.Price = z
		product.Time = time.Now()
		products = append(products, product)
	}
	sqlServer()
}

func sqlServer() {
	PSQLserver := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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

func rows(w http.ResponseWriter, r *http.Request){
	PSQLserver := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", PSQLserver)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected successfully!")
	}
	defer db.Close()
	Rows,err := db.Query("SELECT * FROM infor")
	if err != nil{
		panic(err)
	}
	var rows []Row
	for Rows.Next() {
		var db_row Row
		Rows.Scan(&db_row.ID,&db_row.Item,&db_row.Price,&db_row.Quantity,&db_row.Create_Time)
		rows = append(rows, db_row)
	}
	json.NewEncoder(w).Encode(rows)
}

