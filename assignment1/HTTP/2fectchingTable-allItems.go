package HTTP

import (
	"encoding/json"
	"net/http"

	"github.com/BooMER23/golang.git/DB"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func FetchingallItems(w http.ResponseWriter, r *http.Request) {
	rows := DB.FectchingallItems()				//this will fetch allItems list from the database
	json.NewEncoder(w).Encode(rows)
}