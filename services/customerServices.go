package services

import (
	"github.com/freeddser/pumpkin-rpc/model"
	"github.com/freeddser/pumpkin-rpc/repository"
)

func GetAllCustomers() ([]model.Customer, error) {
	repo := repository.GetCustomersRepository()
	data, err := repo.QueryAllCustomers()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InsertCustomer(id int64, name string, email string, phone string) (model.Customer, error) {
	customer := &model.Customer{ID: id, Name: name, Email: email, Phone: phone}

	repo := repository.GetCustomersRepository()
	err := repo.QueryCustomersInsert(customer)
	if err != nil {
		return model.Customer{}, err
	}
	return *customer, nil
}
