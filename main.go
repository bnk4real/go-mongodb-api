package main

import (
	"go-mongodb-api/config"
	"go-mongodb-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.LoadConfig()
	port := os.Getenv("PORT")
	app := routes.Route(r)
	app.Run(port)
}
