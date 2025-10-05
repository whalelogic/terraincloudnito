package main

import (
	"log"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/whalelogic/templ-go/handlers"
	"github.com/whalelogic/templ-go/models"
	"github.com/whalelogic/templ-go/services"
	"github.com/whalelogic/templ-go/templates"
)

//	func Render(c *fiber.Ctx, component templ.Component) error {
//		c.Set("Content-Type", "text/html")
//		return component.Render(c.Context(), c.Response().BodyWriter())
//	}

var posts []models.Post

func main() {

	// For API access to all posts
	// could be cached in memory at some point
	// TODO add pagination
	// TODO add filtering by tag, category, date, etc.
	// TODO add search
	// Load posts from google cloud storage using API

	allPosts, err := services.GetAllPosts()
	if err != nil {
		log.Fatalf("failed to get posts: %v", err)
	}
	if allPosts != nil {
		log.Printf("Loaded %d posts\n", len(allPosts))
	} else {
		log.Println("No posts found")
	}
	posts = append(posts, allPosts...)
	app := fiber.New(fiber.Config{
		ServerHeader: "WhalerAPI",
		AppName:      "templ-go",
	})

	// Serve static files
	app.Static("/static", "./static")
	app.Static("/public", "./public")

	app.Get("/api", func(c *fiber.Ctx) error {
		data := posts
		return c.JSON(data)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Sort posts by creation date in descending order to get the latest ones first.
		sort.SliceStable(posts, func(i, j int) bool {
			return posts[i].CreatedOn > posts[j].CreatedOn
		})

		// Latest 4 posts.
		latestPosts := posts
		if len(posts) > 4 {
			latestPosts = posts[:4]
		}

		c.Set("Content-Type", "text/html")
		return templates.IndexPage(latestPosts).Render(c.Context(), c.Response().BodyWriter())
	})

	app.Get("/about", handlers.AboutHandler)
	app.Get("/all_posts_page", handlers.HandleAllPosts)
	app.Get("/blog/:slug", handlers.HandlePost)
	app.Get("/resume", handlers.ResumeHandler(models.MyProfile))

	log.Println("Listening on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
