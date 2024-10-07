package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Fenroe/gator/internal/database"
	"github.com/google/uuid"
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
		publishedAt, err := time.Parse(time.RFC1123Z, value.PubDate)
		createPostParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     value.Title,
			Url:       value.Link,
			Description: sql.NullString{
				String: value.Description,
				Valid:  value.Description != "",
			},
			PublishedAt: sql.NullTime{
				Time:  publishedAt,
				Valid: err == nil,
			},
			FeedID: feed.ID,
		}
		state.database.CreatePost(context.Background(), createPostParams)
		fmt.Println("New post added")
	}
	return nil
}
