package service

import (
	"net/http"
	"smartpatrol/app/auth"
	"smartpatrol/app/auth/repository"
	"smartpatrol/models"
	"smartpatrol/pkg/helpers"
	"smartpatrol/pkg/util/response"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB *gorm.DB
	Repository repository.Repository
}

func NewService(DB *gorm.DB, Repository repository.Repository) Service {
	return &serviceImpl{
		DB: DB,
		Repository: Repository,
	}
}

func (s *serviceImpl) AuthenticateUser(user models.SecurityAccount, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (s *serviceImpl) Login(c echo.Context, r auth.LoginRequest) (auth.LoginResponse, error) { 
	var (
		UserRes auth.LoginResponse
	)

	user, err := s.Repository.Login(c, s.DB, r.Email)
	if err != nil {
		return UserRes, response.BuildCustomError(http.StatusUnauthorized,  "Email / Password yang anda masukkan salah!")
	}

	if err := s.AuthenticateUser(user, r.Password); err != nil {
		return UserRes, response.BuildCustomError(http.StatusUnauthorized, "Password yang anda masukkan salah!")
	}

	userID := user.ID
	UserRes.Token, err = helpers.GenerateToken(userID)
	if err != nil {
		return UserRes, response.BuildCustomError(http.StatusInternalServerError, "Failed to generate token")
	}

	return UserRes, nil

	
}