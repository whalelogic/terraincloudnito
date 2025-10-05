package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whalelogic/templ-go/templates"
)

func AboutHandler(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return templates.AboutPage().Render(c.Context(), c.Response().BodyWriter())} 
