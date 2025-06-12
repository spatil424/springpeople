package services

import (
	"restapi/interfaces"
	"restapi/models"
)

type UserService interface {
	GetAllUsers() []models.User
}

type usrService struct {
	repo interfaces.UserRepository
}

func NewUserService(r interfaces.UserRepository) UserService {
	return &usrService{
		repo: r,
	}
}

func (s *usrService) GetAllUsers() []models.User {
	return s.repo.GetAllUsers()
}
