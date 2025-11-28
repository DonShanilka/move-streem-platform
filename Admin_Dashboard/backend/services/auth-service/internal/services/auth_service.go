package services

import (
    "auth-service/internal/models"
    "auth-service/internal/repository"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

func RegisterUser(name, email, password string) error {
    hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    user := models.User{Name: name, Email: email, Password: string(hashed)}
    return repository.CreateUser(user)
}

func LoginUser(email, password string) (string, error) {
    user, err := repository.FindUserByEmail(email)
    if err != nil { return "", err }

    if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
        return "", err
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": user.ID,
        "email": user.Email,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })

    return token.SignedString([]byte("MY_SECRET_KEY"))
}
