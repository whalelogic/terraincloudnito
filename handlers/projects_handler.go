
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"terraincloudnito/templates"
)

func ProjectsHandler(c *fiber.Ctx) error {
    c.Set("Content-Type", "text/html")
	return templates.ProjectsPage().Render(c.Context(), c.Response().BodyWriter())
}
