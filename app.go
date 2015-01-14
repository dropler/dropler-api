package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	m := gin.Default()

	setupRoutes(m)

	port := os.Getenv("PORT")

	m.Run(":" + port)
}
