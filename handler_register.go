package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Fenroe/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(state *state, cmd command) error {
	fmt.Println(cmd.arguments)
	if len(cmd.arguments) < 1 {
		return errors.New("expected a username")
	}
	ctx := context.Background()
	username := cmd.arguments[0]
	queryValues := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}
	newUser, err := state.database.CreateUser(ctx, queryValues)
	if err != nil {
		return fmt.Errorf("error occurred when adding user to database: %v", err)
	}
	state.config.SetUser(username)
	fmt.Printf("User successfully registered: %v\n", newUser)
	return nil
}
