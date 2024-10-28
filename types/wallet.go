package types

type Wallet struct {
	ID       string
	Customer Customer
	Balance  uint64
}
