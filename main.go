package main

import (
	"time"

	"github.com/YeisonKirax/fiber-api-rest/src/config"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: time.RFC822,
		TimeZone:   "America/Santiago",
	}))
	config.InitLogrus()
	// Rutas de usuarios agregada
	usersModule := users.NewUsersModule()
	usersModule.LoadRoutes(app)

	port := config.GetEnvVariable("PORT")
	log.Fatalln(app.Listen(port))
}
