package routers

import (
	"JanusGate/middleware"
	"time"

	fiberUtils "JanusGate/middleware/go-utils/fiber"
	"JanusGate/struct/privates"

	"github.com/gofiber/fiber/v2"
)

func SetupPrivateRoutes(app *fiber.App) {

	// JWT Setup
	app.Use(fiberUtils.AuthenticationMiddleware(fiberUtils.JWTConfig{
		Duration:     15 * time.Minute,
		CookieMaxAge: 15 * 60,
		SetCookies:   true,
		SecretKey:    []byte(middleware.GetEnv("SECRET_KEY")),
	}))

	// Endpoints
	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/private")
	v1Endpoint := publicEndpoint.Group("/v1")

	// Service health check
	v1Endpoint.Get("/", privates.CheckServiceHealth)

}
