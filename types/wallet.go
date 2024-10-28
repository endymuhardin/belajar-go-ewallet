package types

type Wallet struct {
	ID       string
	Customer Customer
	Balance  uint64
}

type WalletStorage interface {
	GetWalletByCustomerId(string) *Wallet
	TopupWallet(w *Wallet, amount uint64)
	Payment(w *Wallet, product string, amount uint64)
	Purchase(w *Wallet, merchant string, remarks string, amount uint64)
}
