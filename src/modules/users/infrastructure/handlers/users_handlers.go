package handlers

import (
	"errors"
	"net/http"

	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/application"

	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/dtos"
	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	usersUseCases *application.UsersUseCases
}

func NewUsersHandler(usersUseCases *application.UsersUseCases) *UsersHandler {

	return &UsersHandler{usersUseCases: usersUseCases}
}

func (uh *UsersHandler) CreateUser(ctx *fiber.Ctx) error {
	var body dtos.CreateUserDTO
	err := ctx.BodyParser(&body)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		errorRes := &fiber.Map{"status": false, "error": err.Error()}
		return ctx.JSON(errorRes)
	}
	if body.Name == "" || body.Surname == "" {
		ctx.Status(http.StatusInternalServerError)
		errorRes := &fiber.Map{"status": false, "error": errors.New(
			"Please specify title and author")}
		return ctx.JSON(errorRes)
	}

	userCreated, userCreatedErr := uh.usersUseCases.ExecuteCreateUser(body)
	if userCreatedErr != nil {
		errorRes := &fiber.Map{"status": false, "error": userCreatedErr.Error()}
		return ctx.JSON(errorRes)
	}

	return ctx.JSON(userCreated)
}

func (uh *UsersHandler) GetUsers(ctx *fiber.Ctx) error {
	users, getUsersError := uh.usersUseCases.ExecuteGetUsers()
	if getUsersError != nil {
		errorRes := &fiber.Map{"status": false, "error": getUsersError.Error()}
		return ctx.JSON(errorRes)
	}
	return ctx.JSON(users)
}

func (uh *UsersHandler) GetUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	user, getUserError := uh.usersUseCases.ExecuteGetUser(userId)
	if getUserError != nil {
		errorRes := &fiber.Map{"status": false, "error": getUserError.Error()}
		return ctx.JSON(errorRes)
	}
	return ctx.JSON(user)
}
