package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pbojar/gator/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {
	// Args check
	if len(cmd.args) > 1 {
		return fmt.Errorf("usage: %s [limit]", cmd.name)
	}

	// Get feed by URL
	var limit int32
	if len(cmd.args) == 1 {
		lim, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("couldn't parse args: %w", err)
		}
		limit = int32(lim)
	} else {
		limit = 2
	}

	// Get posts
	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}
	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found!")
		return nil
	}

	for _, post := range posts {
		printPost(post)
	}
	return nil
}

func printPost(post database.Post) {
	fmt.Println()
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("URL:   %s\n", post.Url)
	fmt.Printf("PubAt: %s\n", post.PublishedAt.Time)
	fmt.Printf("Description: \n\n%s", post.Description)
	fmt.Printf("\n\n============================================================\n")
}
