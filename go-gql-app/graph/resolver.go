package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"go-gql-app/graph/model"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	users []*model.User
	DB    *gorm.DB
}
