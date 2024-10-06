package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Fenroe/gator/internal/config"
	"github.com/Fenroe/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	config   *config.Config
	database *database.Queries
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if _, ok := c.handlers[cmd.name]; !ok {
		return fmt.Errorf("key %s not in commands handlers map", cmd)
	}
	return c.handlers[cmd.name](s, cmd)
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	gatorState := state{
		config:   &cfg,
		database: dbQueries,
	}
	handlers := make(map[string]func(*state, command) error)
	gatorCommands := commands{
		handlers: handlers,
	}
	gatorCommands.register("login", handlerLogin)
	gatorCommands.register("register", handlerRegister)
	gatorCommands.register("reset", handlerReset)
	gatorCommands.register("users", handlerUsers)
	gatorCommands.register("agg", handlerAgg)
	gatorCommands.register("addfeed", handlerAddfeed)
	gatorCommands.register("feeds", handlerFeeds)
	cliArgs := os.Args
	if err != nil {
		fmt.Println(fmt.Errorf("an error occurred while connecting to the database: %v", err))
	}
	if len(cliArgs) < 2 {
		fmt.Println("expected argument")
		os.Exit(1)
	} else {
		cliArgs = cliArgs[1:]
	}
	userCommand := command{
		name:      cliArgs[0],
		arguments: cliArgs[1:],
	}
	err = gatorCommands.run(&gatorState, userCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
