package service

import (
	"smartpatrol/app/company"
	"smartpatrol/app/company/repository"
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
		DB:         DB,
		Repository: Repository,
	}
}

func (s *serviceImpl) GetAll(c echo.Context) ([]company.CompanyResponse, error) {
	var compRes []company.CompanyResponse

	result, err := s.Repository.GetAll(c, s.DB)
	if err != nil {
		return compRes, response.BuildError(response.ErrServerError, err)
	}

	for _, company := range result {
		compRes = append(compRes, company.ToResponse())
	}

	return compRes, nil
}

func (s *serviceImpl) GetById(c echo.Context, id uint) (company.CompanyResponse, error) {
	var compRes company.CompanyResponse

	result, err := s.Repository.GetById(c, s.DB, id)
	if err != nil {
		return compRes, response.BuildError(response.ErrNotFound, err)
	}
	compRes = result.ToResponse()

	return compRes, nil
}