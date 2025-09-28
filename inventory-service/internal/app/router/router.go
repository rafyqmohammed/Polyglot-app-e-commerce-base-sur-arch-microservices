package router

import (
    "inventory-service/internal/app/handlers"
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

    // Health check endpoint
     // Health check endpoint
    r.GET("/health", handlers.HealthCheckHandler)
    r.GET("/info", handlers.HealthCheckHandler)

    r.GET("/products", handlers.GetProductsHandler)
	r.GET("/products/:id", handlers.GetProductByIDHandler)
    r.POST("/products", handlers.CreateProductHandler)
    r.PUT("/products", handlers.UpdateProductHandler)
	r.PUT("/products/:id", handlers.UpdateProductByIDHandler)
    r.DELETE("/products/:id", handlers.DeleteProductHandler)

    return r
}
