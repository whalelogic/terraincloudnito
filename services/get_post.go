package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/whalelogic/templ-go/models"
)

func GetPostBySlug(slug string) (models.Post, error) {
	var posts []models.Post

	// Read the JSON file
	file, err := os.ReadFile("allposts.json")
	if err != nil {
		return models.Post{}, fmt.Errorf("failed to read posts file: %w", err)
	}

	// Unmarshal the JSON data
	if err := json.Unmarshal(file, &posts); err != nil {
		return models.Post{}, fmt.Errorf("failed to unmarshal posts: %w", err)
	}

	// Find the post with the matching slug
	for _, post := range posts {
		if post.Slug == slug {
			return post, nil
		}
	}

	return models.Post{}, fmt.Errorf("post with slug '%s' not found", slug)
}
