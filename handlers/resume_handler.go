// Package handlers provides common handler (mostly render) functions.
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whalelogic/templ-go/templates"
	"github.com/whalelogic/templ-go/models"
)


func ResumeHandler(profile models.Profile) fiber.Handler {
    return func(c *fiber.Ctx) error {
        c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
        // The returned function "closes over" the profile variable,
        // so it has access to the data you passed in.
        return templates.ResumePage(profile).Render(c.Context(), c.Response().BodyWriter())
    }
}
