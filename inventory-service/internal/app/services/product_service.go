package services

import (
    "inventory-service/internal/app/models"
    "inventory-service/internal/app/repositories"
)

func GetProducts() ([]models.Product, error) {
    return repositories.GetAllProducts()
}

func GetProductByID(id int) (models.Product, error) {
    return repositories.GetProductByID(id)
}

func AddProduct(p models.Product) error {
    return repositories.CreateProduct(p)
}

func EditProduct(p models.Product) error {
    return repositories.UpdateProduct(p)
}


func UpdateProductByID(id int, product models.Product) error {
    return repositories.UpdateProductByID(id, product)
}

func RemoveProduct(id int) error {
    return repositories.DeleteProduct(id)
}
