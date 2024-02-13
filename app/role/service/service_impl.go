package service

import (
	"smartpatrol/app/role"
	"smartpatrol/app/role/repository"
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
		Repository: Repository}
}

func (s *serviceImpl) GetAllUser(c echo.Context) ([]role.RoleResponse, error) {
	var roleRes []role.RoleResponse

	result, err := s.Repository.GetAllUser(c, s.DB)
	if err != nil {
		return roleRes, response.BuildError(response.ErrServerError, err)
	}

	for _, role := range result {
		roleRes = append(roleRes, role.ToResponse())
	}

	return roleRes, nil
}