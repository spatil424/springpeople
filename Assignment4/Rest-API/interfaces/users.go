package interfaces

import "rest-api/models"

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(m models.User) (int64, error)
	Login(m models.User) (*models.User, error)
}
