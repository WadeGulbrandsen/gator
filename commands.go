package main

import (
	"context"
	"fmt"

	"github.com/WadeGulbrandsen/gator/internal/database"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	fn, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("Command %q not found", cmd.name)
	}
	return fn(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.User_Name)
		if err != nil {
			return err
		}
		return handler(s, c, user)
	}
}
