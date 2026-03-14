package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	jwtutil "auth-service/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func (s *AuthService) Register(email, password string) (*models.User, error) {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        email,
		PasswordHash: string(hash),
	}

	err = s.Repo.CreateUser(user)

	return user, err
}

func (s *AuthService) Login(email, password string) (string, error) {

	user, err := s.Repo.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", err
	}

	return jwtutil.GenerateToken(user.ID)
}
