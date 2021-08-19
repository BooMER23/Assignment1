package HTTP

import (
	"encoding/json"
	"net/http"

	"github.com/BooMER23/golang.git/DB"
)

func RemainingItems(w http.ResponseWriter, r *http.Request){			//this will fetch the remaining items
	list := DB.Fetchingremainingitems()
	json.NewEncoder(w).Encode(list)
}