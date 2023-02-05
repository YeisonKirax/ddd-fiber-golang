package application

import (
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/dtos"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/interfaces"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/models"
)

type UsersUseCases struct {
	usersDomainService interfaces.IUsersDomainService
}

// Newtype returns new type.
func NewUsersUseCases(usersDomainService interfaces.IUsersDomainService) *UsersUseCases {
	return &UsersUseCases{usersDomainService: usersDomainService}
}

func (uus *UsersUseCases) ExecuteCreateUser(body dtos.CreateUserDTO) (models.User, error) {
	return uus.usersDomainService.CreateNewUser(body)
}

func (uus *UsersUseCases) ExecuteGetUsers() ([]models.User, error) {
	return uus.usersDomainService.GetUsers()
}

func (uus *UsersUseCases) ExecuteGetUser(userId string) (models.User, error) {
	return uus.usersDomainService.GetUser(userId)
}
