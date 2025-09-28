package router

import (
    "customer-service/internal/app/handlers"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Ajoute le middleware CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:4200"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    r.GET("/health", handlers.HealthCheckHandler)
    r.GET("/info", handlers.HealthCheckHandler)
    r.GET("/customers", handlers.GetCustomersHandler)
    r.GET("/customers/:id", handlers.GetCustomerByIDHandler)
    r.POST("/customers", handlers.CreateCustomerHandler)
    r.PUT("/customers/:id", handlers.UpdateCustomerHandler)
    r.DELETE("/customers/:id", handlers.DeleteCustomerHandler)

    return r
}