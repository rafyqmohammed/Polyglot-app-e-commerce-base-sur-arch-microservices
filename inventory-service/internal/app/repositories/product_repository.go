package repositories

import (
    "inventory-service/internal/app/models"
    "inventory-service/pkg/utils"
)

// Récupérer tous les produits
func GetAllProducts() ([]models.Product, error) {
    rows, err := utils.DB.Query("SELECT id, name, quantity, price FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var p models.Product
        if err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price); err != nil {
            return nil, err
        }
        products = append(products, p)
    }
    return products, nil
}

func GetProductByID(id int) (models.Product, error) {
    var product models.Product
    err := utils.DB.QueryRow("SELECT id, name, quantity, price FROM products WHERE id = ?", id).
        Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
    return product, err
}



// Ajouter un produit
func CreateProduct(p models.Product) error {
    _, err := utils.DB.Exec("INSERT INTO products(name, quantity, price) VALUES(?, ?, ?)", p.Name, p.Quantity, p.Price)
    return err
}

// Mettre à jour un produit
func UpdateProduct(p models.Product) error {
    _, err := utils.DB.Exec("UPDATE products SET name=?, quantity=?, price=? WHERE id=?", p.Name, p.Quantity, p.Price, p.ID)
    return err
}

// Mettre a jour un produit par id
func UpdateProductByID(id int, c models.Product) error {
    _, err := utils.DB.Exec("UPDATE products SET name = ?, quantity = ?, price = ? WHERE id = ?",
        c.Name, c.Quantity,c.Price, id)
    return err
}

// Supprimer un produit
func DeleteProduct(id int) error {
    _, err := utils.DB.Exec("DELETE FROM products WHERE id=?", id)
    return err
}
