package types

type Customer struct {
	ID    string
	Name  string
	Email string
	Phone string
}

type CustomerStorage interface {
	GetCustomer(string) *Customer
}
