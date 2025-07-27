package main

import (
	"context"
	"fmt"
	"time"
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
		scrapeFeeds(s)
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
		return fmt.Errorf("error - fetchFeed: %v", err)
	}

	printRSSFeed(rssFeed)
	return nil
}

func printRSSFeed(feed *RSSFeed) {
	fmt.Printf("Title: %s\n", feed.Channel.Title)
	fmt.Printf("Link:  %s\n", feed.Channel.Link)
	fmt.Printf("Desc:  %s\n", feed.Channel.Description)
	fmt.Printf("Items:\n\n")
	for _, item := range feed.Channel.Item {
		fmt.Printf("  * Title: %s\n", item.Title)
		fmt.Printf("  * Link:  %s\n", item.Link)
		fmt.Printf("  * PubD:  %s\n", item.PubDate)
		fmt.Printf("  * Desc:  %s\n", item.Description)
		fmt.Printf("\n==============================\n\n")
	}
}
