package handlers

import (
	"github.com/whalelogic/templ-go/services"
	"github.com/whalelogic/templ-go/templates"
	"github.com/gofiber/fiber/v2"
)



func HandleAllPosts(c *fiber.Ctx) error {
	// 1. Fetch the data using your function.
	posts, err := services.GetAllPosts()
	if err != nil {
		// If there's an error, send a 500 status and the error message.
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// 2. Set the content type.
	c.Set("Content-Type", "text/html")

	// 3. Pass the 'posts' slice to your AllPostsPage component and render it.
	return templates.AllPostsPage(posts).Render(c.Context(), c.Response().BodyWriter())
}
