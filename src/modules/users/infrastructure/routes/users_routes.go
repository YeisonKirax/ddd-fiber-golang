package routes

import (
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/infrastructure/handlers"
	"github.com/gofiber/fiber/v2"
)

type UsersRouter struct {
	usersHandlers *handlers.UsersHandler
}

func NewUsersRouter(usersHandlers *handlers.UsersHandler) *UsersRouter {

	return &UsersRouter{usersHandlers: usersHandlers}
}

func (ur *UsersRouter) LoadRoutes(app fiber.Router) {
	usersRoute := app.Group("/users")
	usersRoute.Post("", ur.usersHandlers.CreateUser)
	usersRoute.Get("", ur.usersHandlers.GetUsers)
	usersRoute.Get("/:id", ur.usersHandlers.GetUser)

}
