package main

import (
	"JanusGate/middleware"
	routers "JanusGate/routers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title JANUS V2
// @version 2.0
// @description JANUS Swagger JanusGate
// @termsOfService http://swagger.io/terms/
// @contact.name FDSAP Support
// @contact.email
// @license.name
// @license.url
// @host 192.168.0.137:5000
// @BasePath /
func main() {
	// Initialize DB here

	// Declare & initialize fiber
	app := fiber.New(fiber.Config{
		UnescapePath: true,
	})

	// Configure application CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Declare & initialize logger
	app.Use(logger.New())

	// Declare & initialize routes
	routers.SetupPublicRoutes(app)
	routers.SetupPrivateRoutes(app)

	// Serve the application
	if middleware.GetEnv("SSL") == "enabled" {
		log.Fatal(app.ListenTLS(
			fmt.Sprintf(":%s", middleware.GetEnv("PORT")),
			middleware.GetEnv("SSL_CERTIFICATE"),
			middleware.GetEnv("SSL_KEY"),
		))
	} else {
		err := app.Listen(fmt.Sprintf(":%s", middleware.GetEnv("PORT")))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
