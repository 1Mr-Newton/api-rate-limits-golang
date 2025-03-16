package main

import (
	"log"
	"time"

	"github.com/1mr-newton/api-rate-limits-golang/config"
	"github.com/1mr-newton/api-rate-limits-golang/handlers"
	"github.com/1mr-newton/api-rate-limits-golang/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		AppName: "API Rate Limiter Demo",
	})

	// Add logger middleware
	app.Use(logger.New())

	// Apply global rate limiter middleware
	app.Use(middleware.NewIPRateLimiter(cfg.GlobalRateLimit, time.Duration(cfg.RateLimitExpiration)*time.Second))

	// Setup routes
	setupRoutes(app, cfg)

	// Start the server
	log.Printf("Server starting on port %d...\n", cfg.Port)
	log.Fatal(app.Listen(cfg.GetAddress()))
}

func setupRoutes(app *fiber.App, cfg *config.Config) {
	// API routes
	api := app.Group("/api")

	// Public endpoints - with global rate limit only
	public := api.Group("/public")
	public.Get("/", handlers.GetPublicInfo)

	// IP endpoint - shows the user's IP address
	api.Get("/ip", handlers.GetIPInfo)

	// User endpoints - with user-specific rate limit
	user := api.Group("/user")
	user.Use(middleware.NewUserRateLimiter(cfg.UserRateLimit, time.Duration(cfg.RateLimitExpiration)*time.Second))
	user.Get("/", handlers.GetUserInfo)

	// Admin endpoints - with admin-specific rate limit
	admin := api.Group("/admin")
	admin.Use(middleware.NewAPIKeyRateLimiter(cfg.AdminRateLimit, time.Duration(cfg.RateLimitExpiration)*time.Second))
	admin.Get("/", handlers.GetAdminInfo)
}
