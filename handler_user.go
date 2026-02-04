package main

import (
	"context"
	"fmt"
	"time"

	"github.com/WadeGulbrandsen/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Usage: %s <name>", cmd.name)
	}
	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      cmd.args[0],
		})
	if err != nil {
		return fmt.Errorf("Couldn't create user: %w", err)
	}
	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return fmt.Errorf("Couldn't set current user: %w", err)
	}
	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Usage: %s <name>", cmd.name)
	}
	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Couldn't find user: %w", err)
	}
	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("Couldn't set current user: %w", err)
	}
	fmt.Printf("Current user set to %q\n", user.Name)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't list users: %w", err)
	}
	for _, user := range users {
		suffix := ""
		if user.Name == s.cfg.User_Name {
			suffix = " (current)"
		}
		fmt.Printf("* %s%s\n", user.Name, suffix)
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
