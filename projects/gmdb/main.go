package main

import (
	"gmdb/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDB()

	r := gin.Default()

	r.Run(":8080")
}
