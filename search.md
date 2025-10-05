# Implementing Search with HTMX

This guide provides the code to implement a search functionality for your blog using HTMX.

## 1. Search Logic

Create a new file `services/search_posts.go` and add the following code. This file will contain the logic for searching posts.

```go
package services

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/whaler/templ-go/models"
)

func SearchPosts(query string) ([]models.Post, error) {
	file, err := os.Open("allposts.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	if err := json.Unmarshal(bytes, &posts); err != nil {
		return nil, err
	}

	if query == "" {
		return posts, nil
	}

	var filteredPosts []models.Post
	for _, post := range posts {
		if strings.Contains(strings.ToLower(post.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(post.Subtitle), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(post.Summary), strings.ToLower(query)) {
			filteredPosts = append(filteredPosts, post)
		}
	}

	return filteredPosts, nil
}
```

## 2. Search Handler

Create a new file `handlers/handle_search.go` and add the following code. This handler will process the search requests.

```go
package handlers

import (
	"net/http"

	"github.com/whaler/templ-go/services"
	"github.com/whaler/templ-go/templates"
)

func HandleSearchPosts(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	posts, err := services.SearchPosts(query)
	if err != nil {
		http.Error(w, "Failed to search posts", http.StatusInternalServerError)
		return
	}

	// We are only rendering the post cards, not the whole page
	component := templates.PostCardList(posts)
	component.Render(r.Context(), w)
}
```

## 3. Add the Search Route

Modify your `main.go` file to add the new `/search` route. You will need to convert your app to use `net/http` handlers, or use an adapter. Given you are using `fiber`, I will show how to adapt the `net/http` handler.

First, you need to get the adaptor by running this command:
`go get github.com/gofiber/adaptor/v2`

Then, in your `main` function, add the following route:

```go
    app.Post("/search", adaptor.HTTPHandlerFunc(handlers.HandleSearchPosts))
```

So your `main.go` will look something like this (showing only the relevant part):

```go
package main

import (
	"log"
	"sort"

	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/adaptor/v2"
	"github.com/whalerapi/templ-go/handlers"
	"github.com/whalerapi/templ-go/models"
	"github.com/whalerapi/templ-go/services"
	"github.com/whalerapi/templ-go/templates"
)

// ... (rest of your main.go file)

func main() {
    // ... (your existing code)

	app.Get("/about", handlers.AboutHandler)
	app.Get("/all_posts_page", handlers.HandleAllPosts)
	app.Get("/blog/:slug", handlers.HandlePost)
	app.Get("/resume", handlers.ResumeHandler(models.MyProfile))
    app.Post("/search", adaptor.HTTPHandlerFunc(handlers.HandleSearchPosts)) // Add this line

	log.Println("Listening on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
```

## 4. Update the Frontend

Modify your `templates/all_posts_page.templ` file to add the search bar with HTMX attributes.

Replace this:
```html
<div class="mb-8">
    <input type="text" placeholder="Search posts..." class="border border-gray-300 rounded-lg py-2 px-4 w-full"/>
</div>
```

With this:
```html
<div class="mb-8">
    <input type="text" name="query" placeholder="Search posts..."
           class="border border-gray-300 rounded-lg py-2 px-4 w-full"
           hx-post="/search"
           hx-trigger="keyup changed delay:500ms"
           hx-target="#search-results"
           hx-swap="innerHTML"
           hx-indicator=".htmx-indicator"
    />
    <span class="htmx-indicator">Searching...</span>
</div>
```

Then, wrap the grid of posts in a `div` with the `id="search-results"`.

Replace this:
```html
<div class="grid grid-cols-1 md:grid-cols-2 gap-8">
    for _, post := range posts {
        @PostCard(post)
    }
</div>
```

With this:
```html
<div id="search-results" class="grid grid-cols-1 md:grid-cols-2 gap-8">
    for _, post := range posts {
        @PostCard(post)
    }
</div>
```

Finally, you need to make sure you have HTMX included in your project. You can add this to your main layout file (`templates/layout.templ`) in the `<head>` section:

```html
<script src="https://unpkg.com/htmx.org@1.9.10"></script>
```

That's it! After making these changes, you will have a working search functionality on your "All Posts" page.
