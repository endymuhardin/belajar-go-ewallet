package types

type Customer struct {
	ID    string
	Name  string `db:"customer_name"`
	Email string
	Phone string `db:"mobile_phone"`
}

type CustomerStorage interface {
	GetCustomer(string) *Customer
}
