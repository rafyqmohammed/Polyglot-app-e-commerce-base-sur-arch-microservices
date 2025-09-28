package repositories

import (
    "customer-service/internal/app/models"
    "customer-service/pkg/utils"
)

func GetAllCustomers() ([]models.Customer, error) {
    rows, err := utils.DB.Query("SELECT id, name, email FROM customers")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var customers []models.Customer
    for rows.Next() {
        var c models.Customer
        if err := rows.Scan(&c.ID, &c.Name, &c.Email); err != nil {
            return nil, err
        }
        customers = append(customers, c)
    }
    return customers, nil
}

func CreateCustomer(c models.Customer) error {
    _, err := utils.DB.Exec("INSERT INTO customers(name, email) VALUES(?, ?)", c.Name, c.Email)
    return err
}

func GetCustomerByID(id int) (models.Customer, error) {
    var customer models.Customer
    err := utils.DB.QueryRow("SELECT id, name, email FROM customers WHERE id = ?", id).
        Scan(&customer.ID, &customer.Name, &customer.Email)
    return customer, err
}

func UpdateCustomer(id int, c models.Customer) error {
    _, err := utils.DB.Exec("UPDATE customers SET name = ?, email = ? WHERE id = ?",
        c.Name, c.Email, id)
    return err
}

func DeleteCustomer(id int) error {
    _, err := utils.DB.Exec("DELETE FROM customers WHERE id = ?", id)
    return err
}