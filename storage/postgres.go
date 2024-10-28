package storage

import "github.com/endymuhardin/belajar-go-ewallet/types"

type PostgresStorage struct{}

func NewPostgresStorage(host string, port string, username string, password string, dbname string) *PostgresStorage {
	return &PostgresStorage{}
}

func (s *PostgresStorage) GetCustomer(id string) *types.Customer {
	return &types.Customer{
		ID:    "abcd",
		Name:  "Customer 99",
		Email: "customer.99@yopmail.com",
		Phone: "081234567890",
	}
}

func (s *PostgresStorage) GetWalletByCustomerId(string) *types.Wallet {
	return nil
}

func (s *PostgresStorage) TopupWallet(w *types.Wallet, amount uint64)                               {}
func (s *PostgresStorage) Payment(w *types.Wallet, product string, amount uint64)                   {}
func (s *PostgresStorage) Purchase(w *types.Wallet, merchant string, remarks string, amount uint64) {}
