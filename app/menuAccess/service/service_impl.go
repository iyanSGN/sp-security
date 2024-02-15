package service

import (
	menuaccess "smartpatrol/app/menuAccess"
	"smartpatrol/app/menuAccess/repository"
	"smartpatrol/pkg/util/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB         *gorm.DB
	Repository repository.Repository
}

func NewService(DB *gorm.DB, Repository repository.Repository) Service {
	return &serviceImpl{
		DB: DB,
		Repository: Repository,
	}
}

func (s *serviceImpl) GetAllMenu(c echo.Context) ([]menuaccess.MenuAccessRes, error) {
	var permissionRes []menuaccess.MenuAccessRes

	result, err := s.Repository.GetAllMenu(c, s.DB)
	if err != nil {
		return permissionRes, response.BuildError(response.ErrServerError, err)
	}

	for _, permission := range result {
		permissionRes = append(permissionRes, permission.ToResponse())
	}

	return permissionRes, nil
}

func (s *serviceImpl) GetMenuByID(c echo.Context, id int) (menuaccess.MenuAccessRes, error) {
	var menuRes menuaccess.MenuAccessRes

	result, err := s.Repository.GetMenuByID(c, s.DB, id)
	if err != nil {
		return menuRes, response.BuildError(response.ErrNotFound, err)
	}

	menuRes = result.ToResponse()

	return menuRes, nil
}

