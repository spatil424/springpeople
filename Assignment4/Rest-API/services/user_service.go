package services

import (
	"rest-api/interfaces"
	"rest-api/models"
	"rest-api/utils"

	"github.com/sirupsen/logrus"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(models.User) (int64, error)
	Login(models.User) (string, error)
}

type userService struct {
	repo interfaces.UserRepository
}

// pass the repository dependency
func NewUserService(r interfaces.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) CreateUser(m models.User) (int64, error) {
	return s.repo.CreateUser(m)
}

func (s *userService) Login(user models.User) (string, error) {
	dbUser, err := s.repo.Login(user)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWTToken(dbUser.ID, dbUser.Username, dbUser.Role)
	if err != nil {
		logrus.Errorf("Error generating JWT token for user %s: %v", dbUser.Username, err)
		return "", err
	}
	return token, nil
}
