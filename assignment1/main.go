package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/BooMER23/golang.git/HTTP"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/allitems", HTTP.AllItems).Methods("POST")				//Enter all the items 
	router.HandleFunc("/fetchitems", HTTP.FetchingallItems).Methods("GET")		//This will fetch all items from the database 
	router.HandleFunc("/sellitems", HTTP.SellItems).Methods("POST")				//Sell the items 
	router.HandleFunc("/customerdetails",HTTP.SoldItems).Methods("GET")			//This will fetch the customer details 
	router.HandleFunc("/rem", HTTP.Itemsremaining).Methods("POST")				//This will insert the remainingitems into database
	router.HandleFunc("/remainingitems",HTTP.RemainingItems).Methods("GET")		//This will fetch the remainingitems from the database 
	http.ListenAndServe(":8080", router)
}