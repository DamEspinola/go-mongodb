package main

import (
	"os"

	"go-gin-image-store/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	router := gin.Default()
	http.Routes(router)
	router.Run(":" + port)
}
