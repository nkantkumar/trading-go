package repository

import "trading-go/internal/model"

type UserRepository interface {
	GetUserByID(id int) (*model.User, error)
	CreateUser(user *model.User) error
	DeleteUser(id int) error
}
