package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/chillyweather/gator/internal/cli"
	"github.com/chillyweather/gator/internal/config"
	"github.com/chillyweather/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)
	appState := cli.State{
		Config: &cfg,
		DB:     dbQueries,
	}

	appCommands := cli.Commands{}
	appCommands.Register("login", cli.HandlerLogin)
	appCommands.Register("register", cli.HandlerRegister)
	appCommands.Register("reset", cli.HandleDelete)
	appCommands.Register("users", cli.HandleGetUsers)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	commandName := args[1]
	commandArguments := args[2:]

	if err := appCommands.Run(&appState, cli.Command{Name: commandName, Args: commandArguments}); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
