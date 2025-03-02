package service

import "user/models"

type ServiceInterface interface {
	CreateUser(user models.User) error
}
