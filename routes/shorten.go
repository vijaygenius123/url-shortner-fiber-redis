package routes

import "C"
import (
	"github.com/gofiber/fiber/v2"
	"time"
	"url-shortner-fiber-redis/helpers"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rateLimit"`
	XRateLimitReset time.Duration `json:"rateLimitReset"`
}

func ShortenUrl(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	// implement rate limiting

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	body.URL = helpers.EnforceHTTP(body.URL)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
