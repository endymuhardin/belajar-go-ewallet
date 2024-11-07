package handler

import (
	"encoding/json"
	"net/http"

	"github.com/endymuhardin/belajar-go-ewallet/types"
)

func HelloHandler(service types.WalletService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		name := r.PathValue("name")
		cust := service.GetCustomer(name)
		json.NewEncoder(w).Encode(cust)
	})
}
