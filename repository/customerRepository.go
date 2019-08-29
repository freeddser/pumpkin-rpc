package repository

import (
	"github.com/freeddser/pumpkin-rpc/model"
)

type CustomersRepository interface {
	QueryCustomersInsert(customer *model.Customer) error
	QueryAllCustomers() ([]model.Customer, error)
}

type customersRepository struct {
	db *DataSource
}

func GetCustomersRepository() CustomersRepository {
	return &customersRepository{db: psqldatasource}
}

func (repo *customersRepository) QueryCustomersInsert(customer *model.Customer) error {
	_, err := repo.db.Exec(`INSERT INTO public.customer(id, name,email,phone) VALUES ($1, $2, $3, $4);`, customer.ID, customer.Name, customer.Email, customer.Phone)
	if err != nil {
		log.Errorf("Error when insert customer: %s", err.Error())
		return err
	}

	return nil
}

func (repo *customersRepository) QueryAllCustomers() ([]model.Customer, error) {
	customers := []model.Customer{}
	err := repo.db.Select(&customers, "SELECT * FROM public.customer")
	return customers, err
}
