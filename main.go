package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// main funge da punto di ingresso per l'applicazione
func main() {
    router := gin.Default()

    // Definire le rotte
    router.GET("/api/hello", helloHandler)

    // Avvia il server sulla porta 8080
    if err := router.Run(":8080"); err != nil {
        panic(err)
    }
}

// helloHandler gestisce la richiesta GET per l'endpoint /api/hello
func helloHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}