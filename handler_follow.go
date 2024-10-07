package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Fenroe/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(state *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return errors.New("expected url argument")
	}
	feed, err := state.database.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}
	currentUserId := user.ID
	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUserId,
		FeedID:    feed.ID,
	}
	follow, err := state.database.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return err
	}
	fmt.Printf("%s is now following %s", follow.UserName, follow.FeedName)
	return nil
}
