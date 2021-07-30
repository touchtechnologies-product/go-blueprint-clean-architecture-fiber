package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
)

func main() {
	// Load config
	appConfig := config.Get()

	// Init log format
	log := setupLog()

	// Gin setup
	router := fiber.New()
	// Set custom log for gin
	router.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}",
	}), cors.New(cors.Config{
		ExposeHeaders: "X-Content-Length",
	}))
	// Jaeger setup
	closer := setupJaeger(appConfig)
	defer func() {
		if err := closer.Close(); err != nil {
			log.Error(err)
		}
	}()

	// Register route to gin
	_ = newApp(appConfig).RegisterRoute(router)

	// Gin start listen
	_ = router.Listen(":8080")
}
