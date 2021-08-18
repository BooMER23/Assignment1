package HTTP

import(
	"net/http"
	"encoding/json"
)
func BuyList(w http.ResponseWriter, r *http.Request) {
	for index := range MulBuy {
		for i := range products {
			if MulBuy[index].Item == products[i].Name {
				Final.Name = MulBuy[index].Name
				Final.Item = MulBuy[index].Item
				Final.Quantity = products[i].Quantity - MulBuy[index].Quantity
				products[i].Quantity = products[i].Quantity - MulBuy[index].Quantity
			}
		}
		UpdatedItems = append(UpdatedItems, Final)
	}
	json.NewEncoder(w).Encode(UpdatedItems)
}
