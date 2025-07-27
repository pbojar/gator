package main

import (
	"context"
	"fmt"
)

func handleAgg(s *state, cmd command) error {
	rssURL := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchFeed(context.Background(), rssURL)
	if err != nil {
		return fmt.Errorf("error - fetchFeed: %v", err)
	}
	fmt.Printf("%+v", rssFeed)
	return nil
}
