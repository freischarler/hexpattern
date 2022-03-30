package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"191", "Martin", "Parana", "3100", "1989-01-01", "1"},
		{"192", "Jose", "Parana", "3100", "1989-02-02", "1"},
	}
	return CustomerRepositoryStub{customers}
}
 