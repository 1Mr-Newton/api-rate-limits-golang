package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetPublicInfo handles the public info endpoint
func GetPublicInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "success",
		"message":   "This is a public endpoint with global rate limiting",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// GetUserInfo handles the user info endpoint
func GetUserInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "success",
		"message":   "This is a user endpoint with user-specific rate limiting",
		"timestamp": time.Now().Format(time.RFC3339),
		"user_id":   c.Query("user_id", "anonymous"),
	})
}

// getClientIP gets the real client IP, prioritizing X-Forwarded-For header
func getClientIP(c *fiber.Ctx) string {
	ips := c.IPs()
	if len(ips) > 0 {
		return ips[0]
	}
	return c.IP()
}

// GetIPInfo handles the IP info endpoint
func GetIPInfo(c *fiber.Ctx) error {
	// Get the client IP that will be used for rate limiting
	clientIP := getClientIP(c)

	// Get the User-Agent
	userAgent := c.Get("User-Agent")

	return c.JSON(fiber.Map{
		"status":         "success",
		"message":        "Your IP address information",
		"timestamp":      time.Now().Format(time.RFC3339),
		"ip":             c.IP(),
		"ips":            c.IPs(),
		"hostname":       c.Hostname(),
		"rate_limit_key": "ip:" + clientIP,
		"client_ip":      clientIP,
		"user_agent":     userAgent,
		"headers": fiber.Map{
			"x_forwarded_for": c.Get("X-Forwarded-For"),
			"x_real_ip":       c.Get("X-Real-IP"),
		},
	})
}

// GetAdminInfo handles the admin info endpoint
func GetAdminInfo(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-Key")
	if apiKey == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "API key is required for admin endpoints",
		})
	}

	return c.JSON(fiber.Map{
		"status":    "success",
		"message":   "This is an admin endpoint with admin-specific rate limiting",
		"timestamp": time.Now().Format(time.RFC3339),
		"api_key":   apiKey[:5] + "...", // Show only first 5 chars for security
	})
}
