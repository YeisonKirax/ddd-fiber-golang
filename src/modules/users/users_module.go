package users

import (
	"github.com/YeisonKirax/fiber-api-rest/src/config"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/application"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/services"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/infrastructure/db/repositories"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/infrastructure/handlers"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/infrastructure/routes"
	"github.com/gofiber/fiber/v2"
)

type UsersModule struct {
}

// Newtype returns new type.
func NewUsersModule() *UsersModule {
	return &UsersModule{}
}

func (*UsersModule) LoadRoutes(app fiber.Router) {
	usersCollection := config.GetCollection(config.DB, "users")
	usersRepository := repositories.NewUsersRepository(usersCollection)
	usersDomainService := services.NewUsersService(usersRepository)
	usersUseCases := application.NewUsersUseCases(usersDomainService)
	usersHandlers := handlers.NewUsersHandler(usersUseCases)
	usersRoutes := routes.NewUsersRouter(usersHandlers)
	usersRoutes.LoadRoutes(app)
}
