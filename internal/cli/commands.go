package cli

import (
	"errors"
	"fmt"

	"github.com/chillyweather/gator/internal/config"
)

type State struct {
	Config *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	command map[string]func(s *State, cmd Command) error
}

func (c Commands) Run(s *State, cmd Command) error {
	handler, ok := c.command[cmd.Name]
	if !ok {
		return errors.New("no such command")
	}

	err := handler(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if c.command == nil {
		c.command = make(map[string]func(s *State, cmd Command) error)

	}
	c.command[name] = f
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("no arguments")
	}

	if err := s.Config.SetUser(cmd.Args[0]); err != nil {
		return fmt.Errorf("error %v", err)
	}
	fmt.Printf("User has been set to %v\n", s.Config.CurrentUserName)

	return nil
}
