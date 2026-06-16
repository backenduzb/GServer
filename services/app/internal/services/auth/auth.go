package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"app/internal/database"
	"app/internal/models"
)

func GetUserByID(userID uint) (models.User, error) {
	var user models.User

	error := database.DB.Preload("Position").First(&user, userID).Error
	if error != nil {
		return user, error
	}

	return user, nil
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

	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}

	return database.DB.Create(&models.Position{
		UserID: user.ID,
	}).Error
}

func GetUser(username string) (models.User, error) {
	var user models.User

	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, errors.New("User not found")
	}

	return user, nil
}