package main

import (
	"context"
	"fmt"
)

func handlerReset(state *state, cmd command) error {
	err := state.database.DeleteAllusers(context.Background())
	if err == nil {
		fmt.Printf("Successfully reset database\n")
	}
	return err
}
