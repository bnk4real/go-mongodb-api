package repositories

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomersRepositories interface {
	GetCustomers(ctx context.Context) ([]Customers, error)
	CreateCustomer(ctx context.Context, customer Customers) error
}

type MongoDBConfig struct {
	Database   string
	Collection string
}

type customersRepositories struct {
	db     *mongo.Client
	config MongoDBConfig
}

func NewRepositories(db *mongo.Client, config MongoDBConfig) CustomersRepositories {
	return &customersRepositories{
		db:     db,
		config: config,
	}
}

type Customers struct {
	CustomerId   string `db:"column:customerid"`
	CustomerName string `db:"column:customername"`
	CustomerType string `db:"column:customertype"`
	Remarks      string `db:"column:remarks"`
}

func (tb *Customers) Collection() string {
	return "customers"
}

// CreateCustomer implements CustomersRepositories.
func (create *customersRepositories) CreateCustomer(ctx context.Context, customer Customers) error {
	database := create.db.Database(create.config.Database)
	collection := database.Collection(create.config.Collection)

	_, err := collection.InsertOne(context.Background(), customer)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// GetCustomers implements CustomersRepositories.
func (find *customersRepositories) GetCustomers(ctx context.Context) ([]Customers, error) {
	filter := bson.M{}

	database := find.db.Database(find.config.Database)
	collection := database.Collection(find.config.Collection)

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var customersData []Customers
	for cursor.Next(context.Background()) {
		var customer Customers
		if err := cursor.Decode(&customer); err != nil {
			log.Fatal(err)
		}
		customersData = append(customersData, customer)
	}

	return customersData, nil
}
