package services

import (
    "customer-service/internal/app/models"
    "customer-service/internal/app/repositories"
)

func GetCustomers() ([]models.Customer, error) {
    return repositories.GetAllCustomers()
}

func AddCustomer(c models.Customer) error {
    return repositories.CreateCustomer(c)
}

func GetCustomerByID(id int) (models.Customer, error) {
    return repositories.GetCustomerByID(id)
}

func UpdateCustomer(id int, customer models.Customer) error {
    return repositories.UpdateCustomer(id, customer)
}

func DeleteCustomer(id int) error {
    return repositories.DeleteCustomer(id)
}