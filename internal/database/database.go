package database

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type service struct {
	db *mongo.Client
}

var (
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	//database = os.Getenv("DB_DATABASE")
)

func New() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)))

	if err != nil {
		log.Fatal(err)

	}
	return client
}
