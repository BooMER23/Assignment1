package HTTP

import (
	"encoding/json"
	"net/http"

	"github.com/BooMER23/golang.git/DB"
)

func SoldItems(w http.ResponseWriter, r *http.Request){			
	details := DB.FetchsoldItems()				//fetching customerdetaisls
	json.NewEncoder(w).Encode(details)
}



