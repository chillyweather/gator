package cli

import (
	"errors"
	"fmt"

	"github.com/chillyweather/gator/internal/config"
)

type State struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	command map[string]func(s *State, cmd command) error
}

func (c commands) run(s *State, cmd command) error {
	handler, ok := c.command[cmd.name]
	if !ok {
		return errors.New("No such command")
	}

	handler(s, cmd)

	return nil
}

func (c *commands) register(name string, f func(*State, command) error) {
	if c.command == nil {
		c.command = make(map[string]func(s *State, cmd command) error)

	}
	c.command[name] = f
}

func handlerLogin(s *State, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("No arguments...")
	}

	if err := s.config.SetUser(cmd.args[0]); err != nil {
		return fmt.Errorf("Error %v", err)
	}
	fmt.Println("User has been set to ")

	return nil
}
