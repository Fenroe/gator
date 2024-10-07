package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Fenroe/gator/internal/database"
)

func handlerBrowse(state *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return errors.New("expected a limit argument")
	}
	limit,err := strconv.Atoi(cmd.arguments[0])
	if err != nil {
		return err
	}
	currentUserId := user.ID
	params := database.GetPostsForUserParams{
		UserID: currentUserId,
		Limit: int32(limit),
	}
	posts, err := state.database.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}
	if len(posts) < 1 {
		fmt.Println("You have no posts saved")
	} else {
		for _,value := range posts {
			fmt.Println(value.Title)
		}
	}
	return nil
}