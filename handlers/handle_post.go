package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whalelogic/templ-go/services"
	"github.com/whalelogic/templ-go/templates"
)

func HandlePost(c *fiber.Ctx) error {
	// 1. Get the slug from the URL.
	slug := c.Params("slug")

	// 2. Fetch the data using your function.
	post, err := services.GetPostBySlug(slug)
	if err != nil {
		// If there's an error, send a 404 status and the error message.
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	// 3. Set the content type.
	c.Set("Content-Type", "text/html")

	// 4. Pass the post to your PostPage component and render it.
	return templates.PostPage(post).Render(c.Context(), c.Response().BodyWriter())
}
