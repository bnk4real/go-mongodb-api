package database

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestNewMongoDB(t *testing.T) {
	_, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected Successfully.")
}
