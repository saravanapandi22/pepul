package services

import "example/pepel/models"

type UserService interface {
	CreateUser(*models.User) error
	GetAll()([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
