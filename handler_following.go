package main

import (
	"context"
	"fmt"

	"github.com/Fenroe/gator/internal/database"
)

func handlerFollowing(state *state, _ command, user database.User) error {
	feedFollows, err := state.database.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Feeds followed by %s:\n", user.Name)
	for _, follow := range feedFollows {
		fmt.Println(follow.FeedName)
	}
	return nil
}
