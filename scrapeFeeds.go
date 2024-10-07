package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Fenroe/gator/internal/database"
)

func scrapeFeeds(state *state) error {
	feed, err := state.database.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	params := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: feed.ID,
	}
	err = state.database.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return err
	}
	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	for _, value := range fetchedFeed.Channel.Item {
		fmt.Printf("%s\n", value.Title)
	}
	return nil
}
