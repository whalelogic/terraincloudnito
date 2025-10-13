// Package handlers provides common handler (mostly render) functions.
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whalelogic/templ-go/models"
	"github.com/whalelogic/templ-go/templates"
)

func ResumeHandler(profile models.Profile) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return templates.ResumePage(profile).Render(c.Context(), c.Response().BodyWriter())
	}
}
