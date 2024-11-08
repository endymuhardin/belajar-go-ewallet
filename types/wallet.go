package types

import (
	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID       string
	Name     string
	Customer Customer
	Balance  decimal.Decimal
}

type WalletTable struct {
	Id            string          `db:"id_wallet"`
	Name          string          `db:"wallet_name"`
	Balance       decimal.Decimal `db:"amount"`
	CustomerId    string          `db:"id_customer"`
	CustomerName  string          `db:"customer_name"`
	CustomerEmail string          `db:"email"`
	CustomerPhone string          `db:"mobile_phone"`
}

type WalletStorage interface {
	GetWalletByCustomerId(string) (*Wallet, error)
	TopupWallet(w *Wallet, amount decimal.Decimal) error
	Payment(w *Wallet, product string, amount decimal.Decimal) error
	Purchase(w *Wallet, merchant string, remarks string, amount decimal.Decimal) error
}

func (walletTable WalletTable) ToWallet() *Wallet {
	cust := Customer{
		ID:    walletTable.CustomerId,
		Name:  walletTable.CustomerName,
		Email: walletTable.CustomerEmail,
		Phone: walletTable.CustomerPhone,
	}

	wallet := Wallet{
		ID:       walletTable.Id,
		Customer: cust,
		Name:     walletTable.Name,
		Balance:  walletTable.Balance,
	}

	return &wallet
}
