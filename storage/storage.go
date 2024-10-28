package storage

import "github.com/endymuhardin/belajar-go-ewallet/types"

type Storage interface {
	GetCustomer(string) *types.Customer
	GetWalletByCustomerId(string) *types.Wallet
	TopupWallet(w *types.Wallet, amount uint64)
	Payment(w *types.Wallet, product string, amount uint64)
	Purchase(w *types.Wallet, merchant string, remarks string, amount uint64)
}
