package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// HelloHandler gestisce la richiesta GET per l'endpoint /api/hello
func HelloHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}