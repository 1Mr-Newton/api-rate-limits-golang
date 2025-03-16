package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// RateLimiterConfig defines the config for the rate limiter middleware
type RateLimiterConfig struct {
	// Max number of requests allowed within the expiration duration
	Max int
	// Expiration is the time on how long to keep records of requests
	Expiration time.Duration
	// KeyGenerator generates the key for each request
	KeyGenerator func(*fiber.Ctx) string
	// LimitReached is called when a request hits the limit
	LimitReached fiber.Handler
	// SkipFailedRequests skips failed requests (status >= 400)
	SkipFailedRequests bool
	// Store is used to store the state of the middleware
	Store fiber.Storage
}

// NewUserRateLimiter creates a new rate limiter middleware for user-specific rate limiting
func NewUserRateLimiter(max int, expiration time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			// Use user_id from query params or header for user-specific rate limiting
			// If not available, fall back to IP address
			userID := c.Query("user_id")
			if userID == "" {
				userID = c.Get("X-User-ID")
			}
			if userID == "" {
				return c.IP()
			}
			return "user:" + userID
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   "Too Many Requests",
				"message": "User rate limit exceeded. Please try again later.",
			})
		},
	})
}

// NewIPRateLimiter creates a new rate limiter middleware based on IP address
func NewIPRateLimiter(max int, expiration time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   "Too Many Requests",
				"message": "IP rate limit exceeded. Please try again later.",
			})
		},
	})
}

// NewAPIKeyRateLimiter creates a new rate limiter middleware based on API key
func NewAPIKeyRateLimiter(max int, expiration time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			apiKey := c.Get("X-API-Key")
			if apiKey == "" {
				// If no API key is provided, use IP as fallback
				return c.IP()
			}
			return "api:" + apiKey
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   "Too Many Requests",
				"message": "API key rate limit exceeded. Please try again later.",
			})
		},
	})
}
