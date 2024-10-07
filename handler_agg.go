package main

import (
	"errors"
	"fmt"
	"time"
)

func handlerAgg(state *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return errors.New("expected a time between reqs argument")
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenReqs)
	fmt.Printf("Fetching feeds every %v\n", timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(state)
	}
}
