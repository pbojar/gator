package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pbojar/gator/internal/database"
)

func handleAddfeed(s *state, cmd command) error {
	// Get current user
	currentUser, err := s.db.GetUser(context.Background(), *s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	// Args check
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]

	// Create entry in feeds
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    currentUser.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	// Create entry in feed_follows
	followCommand := command{
		name: "follow",
		args: []string{feedURL},
	}
	err = handleFollow(s, followCommand)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed, currentUser.Name)
	fmt.Println()
	fmt.Println("Feed followed successfully:")
	printFeedFollow(currentUser.Name, feed.Name)
	fmt.Println("=====================================")
	return nil
}

func handleListFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error - handleListFeeds: %v", err)
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feedRow := range feeds {
		printFeedRow(feedRow)
		fmt.Println("=====================================")
	}
	return nil
}

func printFeed(feed database.Feed, username string) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", username)
}

func printFeedRow(feed database.GetFeedsRow) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", feed.Username.String)
}
