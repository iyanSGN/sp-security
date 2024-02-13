package repository

import (
	"fmt"
	"smartpatrol/app/auth"
	"smartpatrol/models"
	"smartpatrol/pkg/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (*repositoryImpl) Login(c echo.Context, db *gorm.DB, email string) (models.SecurityAccount, error) {
	var user models.SecurityAccount

	if err := db.Where("email = ?", email).First(&user).Error;
		err != nil {
			return user, err
	}

	return user, nil
}

func CreateUser(newUser auth.LoginRequest) (models.SecurityAccount, error) {
	db := database.DBManager()

	user := models.SecurityAccount {
		Email: newUser.Email,
		Password: newUser.Password,
	}

	result := db.Create(&user)
	if result.Error != nil {
		return user, fmt.Errorf("error creating user: %w", result.Error)
	}

	return user, nil
}
