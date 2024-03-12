package repositories

import (
	"context"
	"fmt"
	"go-mongodb-api/database"
	"testing"
)

func TestGetCustomers(t *testing.T) {
	db := database.NewMongoDB()
	config := MongoDBConfig{
		Database:   "customers",
		Collection: "customercl",
	}
	data, err := NewRepositories(db, config).GetCustomers(context.Background())
	if err != nil {
		t.Log(err)
	}

	fmt.Println(data)
}
