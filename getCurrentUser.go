package main

import (
	"context"

	"github.com/Fenroe/gator/internal/database"
)

func getCurrentUser(state *state) (database.User, error) {
	return state.database.GetUser(context.Background(), state.config.CurrentUserName)
}
