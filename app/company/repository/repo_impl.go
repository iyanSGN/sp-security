package repository

import (
	"fmt"
	"smartpatrol/app/company"
	"smartpatrol/models"
	"smartpatrol/pkg/database"
	"smartpatrol/pkg/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll(c echo.Context, DB *gorm.DB) ([]models.CompanyCompany, error) {
	var comp []models.CompanyCompany

	if err := DB.Find(&comp).Error; err != nil {
		return nil, err
	}

	return comp, nil
}

func (r *repositoryImpl) GetById(c echo.Context, db *gorm.DB, id uint) (models.CompanyCompany, error) {
	var comp models.CompanyCompany

	// Include the CompanyBuilding association to fetch related buildings
	if err := db.Where("id = ?", id).First(&comp).Error; err != nil {
		return comp, err
	}

	return comp, nil
}

func CreateCompany(req company.CompanyRequest) (company.CompanyRequest, error) {
	db := database.DBManager()

	comp := models.CompanyCompany{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy: req.CreatedBy,
	}

	result := db.Create(&comp)
	if result.Error != nil {
		return company.CompanyRequest{}, fmt.Errorf("error creating company: %w", result.Error)
	}

	return req, nil
}

func UpdateCompany(id int, update company.CompanyRequest) error {
	db := database.DBManager()

	var comp models.CompanyCompany

	result := db.First(&comp, id)
	if result.Error != nil {
		return fmt.Errorf("error retrieving existing company: %w", result.Error)
	}

	if update.Name != "" {
		comp.Name = update.Name
	}

	if update.UpdatedBy != 0 {
		comp.UpdatedBy = update.UpdatedBy
	}

	if update.Description != "" {
		comp.Description = update.Description
	}

	UpdatedCompany := db.Save(&comp)
	if UpdatedCompany.Error != nil {
		return fmt.Errorf("error updating company: %w", UpdatedCompany.Error)
	}

	return nil

}

func DeleteCompany(id int) error {
	db := database.DBManager()

	var compID models.CompanyCompany
	result := db.First(&compID, id)
	if result.Error != nil {
		return fmt.Errorf("error retrieving company: %w", result.Error)
	}

	deletedResult := db.Delete(&compID)
	if deletedResult.Error != nil {
		return fmt.Errorf("error deleting company: %w", deletedResult.Error)
	}

	return nil
}

func IsActiveCompany(id uint, company *models.CompanyCompany) error {
	db := database.DBManager()

	result := db.First(&company, id)
	if result.Error != nil {
		return fmt.Errorf("error retrieving company: %w", result.Error)
	}

	if err := helpers.ChecklistIsActive(&company.IsActive); err != nil {
		return fmt.Errorf("error checklist/unchecklist IsActive: %w", err)
	}

	if err := db.Save(&company).Error; err != nil {
		return fmt.Errorf("error saving company: %w", err)
	}

	return nil
}