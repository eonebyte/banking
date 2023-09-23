package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Iwan", City: "Cikarang", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "John", City: "Bogor", Zipcode: "220011", DateofBirth: "2000-06-02", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
