package main

import _ "github.com/lib/pq"

import (
	"fmt"
	"os"

	"github.com/chillyweather/gator/internal/cli"
	"github.com/chillyweather/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	appState := cli.State{
		Config: &cfg,
	}

	appCommands := cli.Commands{}
	appCommands.Register("login", cli.HandlerLogin)

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
