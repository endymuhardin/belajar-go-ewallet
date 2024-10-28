package storage

import "github.com/endymuhardin/belajar-go-ewallet/types"

type Storage interface {
	types.CustomerStorage
	types.WalletStorage
}
