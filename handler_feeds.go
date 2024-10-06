package main

import (
	"context"
	"fmt"
)

func handlerFeeds(state *state, _ command) error {
	feeds, err := state.database.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("FEED NAME: %s; FEED URL: %s, USER NAME: %s\n", feed.FeedName, feed.FeedUrl, feed.UserName)
	}
	return nil
}
