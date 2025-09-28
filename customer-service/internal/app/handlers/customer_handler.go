package handlers

import (
    "customer-service/internal/app/models"
    "customer-service/internal/app/services"
    "github.com/gin-gonic/gin"
    "net/http"
	"strconv"
)

func GetCustomersHandler(c *gin.Context) {
    customers, err := services.GetCustomers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, customers)
}

func CreateCustomerHandler(c *gin.Context) {
    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if err := services.AddCustomer(customer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Customer created"})
}

func GetCustomerByIDHandler(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    customer, err := services.GetCustomerByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
        return
    }
    c.JSON(http.StatusOK, customer)
}

func UpdateCustomerHandler(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if err := services.UpdateCustomer(id, customer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Customer updated"})
}
func DeleteCustomerHandler(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    if err := services.DeleteCustomer(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
// HealthCheckHandler provides a simple health check endpoint
func HealthCheckHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "UP",
        "service": "customer-service",
    })
}