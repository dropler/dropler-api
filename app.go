package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	m := gin.Default()

	setupDb()
	setupRoutes(m)

	port := os.Getenv("PORT")

	m.Run(":" + port)
}
