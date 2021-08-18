package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/BooMER23/golang.git/HTTP"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HTTP.HomePage).Methods("POST")
	router.HandleFunc("/rows", HTTP.Rows).Methods("GET")
	router.HandleFunc("/buy", HTTP.BuyItems).Methods("POST")
	router.HandleFunc("/buylist", HTTP.BuyList).Methods("GET")
	router.HandleFunc("/finallist",HTTP.FinalList).Methods("POST")
	router.HandleFunc("/data", HTTP.Data).Methods("GET")
	http.ListenAndServe(":8080", router)
}