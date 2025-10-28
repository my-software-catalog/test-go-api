package app

import (
    "github.com/gin-gonic/gin"
    "your-api/internal/routes"
)

// Run avvia il server API
func Run() {
    r := gin.Default()
    routes.SetupRoutes(r)

    if err := r.Run(":8080"); err != nil {
        panic(err)
    }
}