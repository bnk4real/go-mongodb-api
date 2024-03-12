package routes

import (
	"go-mongodb-api/controllers"
	"go-mongodb-api/database"
	"go-mongodb-api/repositories"
	"go-mongodb-api/services"
	"os"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) *gin.Engine {

	db := database.NewMongoDB()
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_COLLECTION")
	config := repositories.MongoDBConfig{
		Database:   dbHost,
		Collection: dbName,
	}

	repo := repositories.NewRepositories(db, config)
	srv := services.NewServicex(repo)
	ctrl := controllers.NewCustomerController(srv)

	// middle ware
	v1 := router.Group("/api/mongodb/go")

	v1.GET("/getxxx", ctrl.GetXXX)
	v1.GET("/get-customers", ctrl.GetCustomer)
	v1.POST("/create-customer", ctrl.CreateCustomer)

	return router
}
