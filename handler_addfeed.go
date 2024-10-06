package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Fenroe/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddfeed(state *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return errors.New("expected name and url")
	}
	name, url := cmd.arguments[0], cmd.arguments[1]
	currentUser, err := getCurrentUser(state)
	if err != nil {
		return err
	}
	currentUserId := currentUser.ID
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
	fmt.Printf("Successfully added new feed: %v\n", feed)
	return nil
}
