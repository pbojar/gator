package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pbojar/gator/internal/database"
)

func handleFollow(s *state, cmd command, user database.User) error {
	// Args check
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
	}

	// Get feed by URL
	feedURL := cmd.args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	// Create entry in feed_follow
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	// Print sucess message
	fmt.Println("Feed follow created:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handleFollowing(s *state, cmd command) error {
	// Args check
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), *s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", *s.cfg.CurrentUserName)
	for _, ff := range feeds {
		fmt.Printf("* %s\n", ff.FeedName.String)
	}

	return nil
}

func handleUnfollow(s *state, cmd command, user database.User) error {
	// Args check
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
	}
	feedURL := cmd.args[0]

	deleteFeed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	for _, feedFollow := range feedFollows {
		if feedFollow.FeedID == deleteFeed.ID {
			err = s.db.DeleteFeedFollow(context.Background(), feedFollow.ID)
			if err != nil {
				return fmt.Errorf("couldn't delete feed follow: %w", err)
			}
			return nil
		}
	}
	fmt.Println("Current user is not following that feed.")
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
