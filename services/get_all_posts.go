package services

// Package /services provides helpful services that are NOT handlers
// They passs data TO the handlers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/whalelogic/templ-go/models"
)

func GetAllPosts() ([]models.Post, error) {

	f, err := os.ReadFile("allposts.json")
	if err != nil {
		return nil, fmt.Errorf("error reading allposts.json: %w", err)
	}

	var posts []models.Post

	if err := json.Unmarshal(f, &posts); err != nil {
		return nil, fmt.Errorf("failed to unmarshall posts: %w", err)
	}

	return posts, nil

}
