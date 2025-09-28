package handlers

import (
    "inventory-service/internal/app/models"
    "inventory-service/internal/app/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func GetProductsHandler(c *gin.Context) {
    products, err := services.GetProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
}

func GetProductByIDHandler(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    product, err := services.GetProductByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}



func CreateProductHandler(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if err := services.AddProduct(product); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

func UpdateProductHandler(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if err := services.EditProduct(product); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

func UpdateProductByIDHandler(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if err := services.UpdateProductByID(id, product); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "product updated"})
}

func DeleteProductHandler(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := services.RemoveProduct(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

// HealthCheckHandler provides a simple health check endpoint
func HealthCheckHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "UP",
        "service": "inventory-service",
    })
}
