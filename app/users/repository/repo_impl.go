package repository

import (
	"fmt"
	"smartpatrol/app/users"
	"smartpatrol/models"
	"smartpatrol/pkg/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct{}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (*repositoryImpl) GetAllUser(c echo.Context, db *gorm.DB) ([]models.SecurityUser, error) {
	var user []models.SecurityUser

	if err := db.Find(&user); err != nil {
		return user, nil
	}

	return user, nil
}

func (*repositoryImpl) GetUserByID(c echo.Context, db *gorm.DB, id int) (models.SecurityUser, error) {
	var role models.SecurityUser

	if err := db.Where("id", id).
		Find(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}

func CreateUser(new users.UserRequest) (users.UserRequest, error) {
	db := database.DBManager()

	newUser := models.SecurityUser{
		CreatedBy:      new.CreatedBy,
		IsActive:       new.IsActive,
		AccountID:      new.AccountID,
		RoleID:         new.RoleID,
		CompanyAreaID:  new.CompanyAreaID,
		CompanyShiftID: new.CompanyShiftID,
		EmployeeID:     new.EmployeeID,
		Name:           new.Name,
	}

	result := db.Create(&newUser)
	if result.Error != nil {
		return users.UserRequest{}, fmt.Errorf("error creating user: %w", result.Error)
	}

	newAccount := models.SecurityAccount{
		ID: newUser.ID,
		CreatedBy: new.CreatedBy,
		IsActive: new.IsActive,
		Email: new.Email,
		Password: new.Password,
	}

	resultAccount := db.Create(&newAccount)
	if resultAccount.Error != nil {
		return new, fmt.Errorf("error creating account: %w", resultAccount.Error)
	}

	return new, nil
}

func UpdateUser(id int, update users.UserRequest) error {
	db := database.DBManager()

	var users models.SecurityUser
	result := db.First(&users, id)

	if result.Error != nil {
		return fmt.Errorf("error retrieving existing user: %w", result.Error)
	}

	if update.Name != "" {
		users.Name = update.Name
	}

	if update.IsActive != 0 {
		users.IsActive = update.IsActive
	}

	if update.AccountID != 0 {
		users.AccountID = update.AccountID
	}

	if update.RoleID != 0 {
		users.RoleID = update.RoleID
	}

	if update.CompanyAreaID != 0 {
		users.CompanyAreaID = update.CompanyShiftID
	}

	if update.EmployeeID != "" {
		users.EmployeeID = update.EmployeeID
	}

	updatedUser := db.Save(&users)
	if updatedUser.Error != nil {
		return fmt.Errorf("error updating User: %w", updatedUser.Error)
	}

	return nil

}


func DeleteUser(id int) error {
	db := database.DBManager()

	var userID models.SecurityUser
	result := db.First(&userID, id)
	if result.Error != nil {
		return fmt.Errorf("error deleting selected user: %w", result.Error)
	}

	deleteUser := db.Delete(&userID)
	if deleteUser.Error != nil {
		return fmt.Errorf("error deleting selected user: %w", deleteUser.Error)
	}

	return nil
}