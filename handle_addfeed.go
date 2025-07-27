package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pbojar/gator/internal/database"
)

func handleAddfeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("error - handleAddfeed: expected 2 arguments")
	}
	feedName := cmd.args[0]
	feedURL := cmd.args[1]
	currentUser, err := s.db.GetUser(context.Background(), *s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error - handleAddfeed: %v", err)
	}
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    currentUser.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error - database.CreateUser: %v", err)
	}
	fmt.Printf("%+v", feed)
	return nil
}
