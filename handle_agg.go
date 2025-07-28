package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pbojar/gator/internal/database"
)

func handleAgg(s *state, cmd command) error {
	// Args check
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_requests>", cmd.name)
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse duration: %w", err)
	}
	fmt.Printf("Collecting feeds every %s\n", timeBetweenReqs.String())

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("couldn't scrape feed: %w", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	// Get next feed to fetch
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get next feed to fetch: %w", err)
	}

	fmt.Printf("Fetching feed: %s...\n", feedToFetch.Name)

	// Mark feed as fetched
	err = s.db.MarkFeedFetched(context.Background(), feedToFetch.ID)
	if err != nil {
		fmt.Printf("%v", err)
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}

	// Fetch feed
	rssFeed, err := fetchFeed(context.Background(), feedToFetch.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	// Add items to posts
	for _, item := range rssFeed.Channel.Item {
		_, err := s.db.GetPostByURL(context.Background(), item.Link)
		if err == nil {
			continue
		}
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			FeedID:      feedToFetch.ID,
		}
		loc, _ := time.LoadLocation("UTC")
		pubDate, err := time.ParseInLocation(time.RFC1123Z, item.PubDate, loc)
		if err == nil {
			params.PublishedAt = sql.NullTime{Valid: true, Time: pubDate}
		}
		_, err = s.db.CreatePost(context.Background(), params)
		if err != nil {
			return fmt.Errorf("couldn't create post: %w", err)
		}
	}

	return nil
}
