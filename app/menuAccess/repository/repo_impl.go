package repository

import (
	"fmt"
	menuaccess "smartpatrol/app/menuAccess"
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

func (r *repositoryImpl) GetAllMenu(c echo.Context, DB *gorm.DB) ([]models.SecurityPermission, error) {
	var permission []models.SecurityPermission

	if err := DB.Find(&permission).Error; err != nil {
		return nil, err
	}

	return permission, nil
}

func (r *repositoryImpl) GetMenuByID(c echo.Context, DB *gorm.DB, id int) (models.SecurityPermission, error) {
	var permission models.SecurityPermission

	if err := DB.Where("id = ?", id).First(&permission).Error; err != nil {
		return permission, err
	}

	return permission, nil
}




func CreateMenuAccess(req menuaccess.MenuAccessReq ) (menuaccess.MenuAccessReq, error) {
	db := database.DBManager()

	permission := models.SecurityPermission{
		CreatedBy: req.CreatedBy,
		IsActive:    req.IsActive,
		Code: req.Code,
		Name:        req.Name,
		Description:   req.Description,
	}

	result := db.Create(&permission)
	if result.Error != nil {
		return req, fmt.Errorf("error updating menu access: %w", result.Error)
	}

	return req, nil
}

func UpdateMenuAccess(id int, req menuaccess.MenuAccessReq) error {
	db := database.DBManager()

	var menu models.SecurityPermission

	result := db.First(&menu, id)
	if result.Error != nil {
		return fmt.Errorf("error updating menu access: %w", result.Error)
	}

	if req.UpdatedBy != 0 {
		menu.UpdatedBy = req.UpdatedBy
	}

	if req.IsActive != 0 {
		menu.IsActive = req.IsActive
	}

	if req.Name != "" {
		menu.Name = req.Name
	}

	if req.Code != "" {
		menu.Code = req.Code
	}

	if req.Description != "" {
		menu.Description = req.Description
	}

	updatedDepartment := db.Save(&menu)
	if updatedDepartment.Error != nil {
		return fmt.Errorf("error updating menu access: %w", updatedDepartment.Error)
	}

	return nil
}

func DeleteDepartment(id int) error {
	db := database.DBManager()

	var menu models.SecurityPermission

	result := db.First(&menu, id)
	if result.Error != nil {
		return fmt.Errorf("error retrieving menu access: %w", result.Error)
	}

	deleteMenu := db.Delete(&menu)
	if deleteMenu.Error != nil {
		return fmt.Errorf("error deleting menu access: %w", deleteMenu.Error)
	}

	return nil
}