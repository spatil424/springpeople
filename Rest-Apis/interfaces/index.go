package interfaces

import "restapi/models"

type UserRepository interface {
	GetAllUsers() []models.User
	GetUserById(id string) (models.User, error)
}
