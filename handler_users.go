package main

import (
	"context"
	"fmt"
)

func handlerUsers(state *state, _ command) error {
	users, err := state.database.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user == state.config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}
	return nil
}
