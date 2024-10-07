package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fenroe/gator/internal/database"
)

func handlerUnfollow(state *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return errors.New("expected a url argument")
	}
	feedUnfollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    cmd.arguments[0],
	}
	err := state.database.DeleteFeedFollow(context.Background(), feedUnfollowParams)
	if err != nil {
		return err
	}
	fmt.Printf("%s has unfollowed %s\n", user.Name, cmd.arguments[0])
	return nil
}
