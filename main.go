package main

import (
	"github.com/avilldaniel/pullup-api/config"
	"github.com/avilldaniel/pullup-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run("localhost:8080")
}
