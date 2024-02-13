package helpers

import (
	"smartpatrol/app/auth"
	"smartpatrol/app/users"
	// "smartpatrol/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(user *auth.LoginRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}

func HashedPassword(user *users.UserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}