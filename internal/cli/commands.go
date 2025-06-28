package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/chillyweather/gator/internal/config"
	"github.com/chillyweather/gator/internal/database"
	"github.com/google/uuid"
)

type State struct {
	Config *config.Config
	DB     *database.Queries
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

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("no arguments")
	}
	name := cmd.Args[0]

	_, err := s.DB.GetUser(context.Background(), name)
	if err == nil {
		os.Exit(1)
		return fmt.Errorf("user '%s' already exists", name)
	}

	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("could not create user: %v", err)
	}

	s.Config.SetUser(user.Name)

	fmt.Printf("User '%s' created successfully!\n", user.Name)
	return nil
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("no arguments")
	}

	_, err := s.DB.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		os.Exit(1)
		return fmt.Errorf("user '%s' does not exist", cmd.Args[0])
	}

	if err := s.Config.SetUser(cmd.Args[0]); err != nil {
		return fmt.Errorf("error %v", err)
	}
	fmt.Printf("User has been set to %v\n", s.Config.CurrentUserName)

	return nil
}
