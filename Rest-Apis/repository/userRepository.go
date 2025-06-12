package repository

import (
	"restapi/interfaces"
	"restapi/models"
)

// implement the UserRepository interface
type UserRepositoryImpl struct{}

func (r *UserRepositoryImpl) GetAllUsers() []models.User {
	return []models.User{
		{Id: "1", Name: "Alice"},
		{Id: "2", Name: "Bob"},
	}
}

func (r *UserRepositoryImpl) GetUserById(id string) (models.User, error) {
	return models.User{Id: id, Name: "John"}, nil
}

func NewRepostiory() interfaces.UserRepository {
	return &UserRepositoryImpl{}
}
