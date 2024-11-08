package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/endymuhardin/belajar-go-ewallet/types"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type PostgresStorage struct {
	dbpool *pgxpool.Pool
}

func NewPostgresStorage(host string, port string, username string, password string, dbname string) *PostgresStorage {
	var url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbname)
	log.Default().Printf("Connecting to %s", url)

	dbpool, err := pgxpool.New(context.Background(), url)
	dbpool.Config().AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdecimal.Register(conn.TypeMap())
		return nil
	}

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
		log.Fatalf("query failed: %s", err)
		return nil, err
	}

	customer, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[types.Customer])

	if err != nil {
		log.Fatalf("unable to convert rows to struct: %s", err)
		return nil, err
	}

	return &customer, nil
}

func (s *PostgresStorage) GetWalletByCustomerId(id string) (*types.Wallet, error) {
	query := `
	select w.id as id_wallet, w.wallet_name, w.amount, 
	c.id as id_customer, c.customer_name, c.email, c.mobile_phone  
	from wallet w inner join customer c 
	on w.id_customer = c.id 
	where w.id_customer = @idCustomer
	`

	args := pgx.NamedArgs{
		"idCustomer": id,
	}

	rows, err := s.dbpool.Query(context.Background(), query, args)

	if err != nil {
		log.Fatalf("query failed: %s", err)
		return nil, err
	}

	walletTable, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[types.WalletTable])

	if err != nil {
		log.Fatalf("unable to convert rows to struct: %s", err)
		return nil, err
	}

	log.Printf("Balance : %s", walletTable.Balance.String())

	return walletTable.ToWallet(), nil
}

func (s *PostgresStorage) TopupWallet(w *types.Wallet, amount decimal.Decimal) error {
	return nil
}

func (s *PostgresStorage) Payment(w *types.Wallet, product string, amount decimal.Decimal) error {
	return nil
}

func (s *PostgresStorage) Purchase(w *types.Wallet, merchant string, remarks string, amount decimal.Decimal) error {
	return nil
}
