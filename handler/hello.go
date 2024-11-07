package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/endymuhardin/belajar-go-ewallet/types"
)

func HelloHandler(service types.WalletService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		name := r.PathValue("name")
		cust, err := service.GetCustomer(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running query: %v\n", err)
			os.Exit(1)
		}
		json.NewEncoder(w).Encode(cust)
	})
}
