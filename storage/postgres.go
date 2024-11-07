package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/endymuhardin/belajar-go-ewallet/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	dbpool *pgxpool.Pool
}

func NewPostgresStorage(host string, port string, username string, password string, dbname string) *PostgresStorage {
	var url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbname)
	log.Default().Printf("Connecting to %s", url)
	dbpool, err := pgxpool.New(context.Background(), url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return &PostgresStorage{dbpool: dbpool}
}

func (s *PostgresStorage) Shutdown() {
	defer s.dbpool.Close()
}

func (s *PostgresStorage) GetCustomer(id string) (*types.Customer, error) {
	query := "select * from customer where id = @id"
	args := pgx.NamedArgs{
		"id": id,
	}
	rows, err := s.dbpool.Query(context.Background(), query, args)

	if err != nil {
		log.Fatalf("unable to query users: %s", err)
		return nil, err
	}

	customer, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[types.Customer])

	if err != nil {
		log.Fatalf("unable to get rows to struct: %s", err)
		return nil, err
	}

	return &customer, nil
}

func (s *PostgresStorage) GetWalletByCustomerId(string) *types.Wallet {
	return nil
}

func (s *PostgresStorage) TopupWallet(w *types.Wallet, amount uint64)                               {}
func (s *PostgresStorage) Payment(w *types.Wallet, product string, amount uint64)                   {}
func (s *PostgresStorage) Purchase(w *types.Wallet, merchant string, remarks string, amount uint64) {}
