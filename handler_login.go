package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(state *state, cmd command) error {
	fmt.Println(cmd.arguments)
	if len(cmd.arguments) < 1 {
		return errors.New("expected a username")
	}
	username := cmd.arguments[0]
	_,err := state.database.GetUser(context.Background(),username)
	if err != nil {
		return errors.New("user not in database")
	}
	state.config.SetUser(username)
	fmt.Printf("You are now logged in as %s\n", username)
	return nil
}
