package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/endymuhardin/belajar-go-ewallet/types"
)

func GetWalletByCustomerId(service types.WalletService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customer := r.PathValue("id")
		log.Printf("Customer ID : %s", customer)
		wallet, err := service.GetWalletByCustomerId(customer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running query: %v\n", err)
			os.Exit(1)
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(wallet)
	})
}
