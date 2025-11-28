package repository

import (
	"context"
	"github.com/DonShanilka/auth-service/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var userCollection *mongo.Collection

func init() {
	userCollection = database.GetDatabase("movie_db").Collection("users")
}

func CreateUser(name, email, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, bson.M{
		"name":     name,
		"email":    email,
		"password": password,
	})
	return err
}
