package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"app/internal/database"
	"app/internal/models"
	"app/config/settings"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GetUserByID(userID uint) (models.User, error) {
	var user models.User

	result := database.DB.First(&user, userID)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func GenerateJWT(userID uint) (string, error) {
	claims :=  jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(settings.Envs.JWT_SECRET))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(username, password string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: hash,
	}
	return database.DB.Create(&user).Error
}

func GetUser(username string) (models.User, error) {
	var user models.User

	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, errors.New("User not found")
	}

	return user, nil
}