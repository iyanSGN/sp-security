package repository

import (
	"smartpatrol/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAllUser(c echo.Context, DB *gorm.DB) ([]models.SecurityRole, error) {
	var department []models.SecurityRole

	if err := DB.Find(&department).Error; err != nil {
		return nil, err
	}

	return department, nil
}