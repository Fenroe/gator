package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Fenroe/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddfeed(state *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return errors.New("expected name and url")
	}
	name, url := cmd.arguments[0], cmd.arguments[1]
	currentUserId := user.ID
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    currentUserId,
	}
	feed, err := state.database.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}
	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUserId,
		FeedID:    feed.ID,
	}
	_, err = state.database.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully added new feed: %v\n", feed)
	return nil
}
