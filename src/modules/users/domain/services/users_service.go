package services

import (
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/dtos"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/interfaces"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/models"
)

type UsersDomainService struct {
	usersRepository interfaces.IUsersRepository
}

func NewUsersService(usersRepository interfaces.IUsersRepository) *UsersDomainService {
	return &UsersDomainService{usersRepository: usersRepository}
}

func (us *UsersDomainService) CreateNewUser(newUser dtos.CreateUserDTO) (models.User, error) {
	return us.usersRepository.Create(newUser)
}

func (us *UsersDomainService) GetUsers() ([]models.User, error) {
	return us.usersRepository.Find()
}

func (us *UsersDomainService) GetUser(userId string) (models.User, error) {
	return us.usersRepository.FindById(userId)
}
