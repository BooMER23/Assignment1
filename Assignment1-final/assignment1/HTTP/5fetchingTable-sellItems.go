package HTTP

import (
	"net/http"

	"github.com/BooMER23/golang.git/DB"
)

type remitem struct {
	Item     string  `json:"Item"`
	Quantity float64 `json:"Quantity"`
}

var Final remitem

var UpdatedItems []remitem

func Itemsremaining(w http.ResponseWriter, r *http.Request) {
	DB.Insertintoremainingitems()
}