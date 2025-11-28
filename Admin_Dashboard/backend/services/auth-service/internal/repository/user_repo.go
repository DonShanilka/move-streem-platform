package repository

import (
    "auth-service/internal/config"
    "auth-service/internal/models"
    "context"
)

func CreateUser(user models.User) error {
    _, err := config.DB.Database("movieapp").Collection("users").InsertOne(context.TODO(), user)
    return err
}

func FindUserByEmail(email string) (models.User, error) {
    var user models.User
    err := config.DB.Database("movieapp").Collection("users").FindOne(
        context.TODO(),
        map[string]interface{}{"email": email},
    ).Decode(&user)
    return user, err
}
