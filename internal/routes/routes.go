package routes

import (
    "github.com/gin-gonic/gin"
    "your-api/internal/handlers"
)

// SetupRoutes definisce le rotte per l'applicazione
func SetupRoutes(r *gin.Engine) {
    r.GET("/api/hello", handlers.HelloHandler)
}