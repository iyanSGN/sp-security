package service

import (
	"smartpatrol/app/users"
	"smartpatrol/app/users/repository"
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

func (s *serviceImpl) GetAllUser(c echo.Context) ([]users.UserResponse, error) {
	var resApproval []users.UserResponse

	result, err := s.Repository.GetAllUser(c, s.DB)
	if err != nil {
		return resApproval, response.BuildError(response.ErrNotFound, err)
	}

	for _, employee := range result {
		resApproval = append(resApproval, employee.ToResponse())
	}

	return resApproval, nil
}

func (s *serviceImpl) GetUserByID(c echo.Context, id int) (users.UserResponse, error) {
	var userRes users.UserResponse

	result, err := s.Repository.GetUserByID(c, s.DB, int(id))
	if err != nil {
		return userRes, response.BuildError(response.ErrNotFound, err)
	}

	userRes = result.ToResponse()

	return userRes, nil
}
