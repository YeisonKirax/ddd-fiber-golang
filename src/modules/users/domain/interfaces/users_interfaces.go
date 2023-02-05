package interfaces

import (
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/dtos"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/models"
)

type IUsersDomainService interface {
	CreateNewUser(newUser dtos.CreateUserDTO) (models.User, error)

	GetUsers() ([]models.User, error)

	GetUser(userId string) (models.User, error)
}

type IUsersRepository interface {
	Create(newUser dtos.CreateUserDTO) (models.User, error)

	Find() ([]models.User, error)

	FindById(userId string) (models.User, error)
}
