package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pbojar/gator/internal/database"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <user_name>", cmd.name)
	}
	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return fmt.Errorf("'%s' is not a registered user", username)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}
	fmt.Printf("Successfully switched user to '%s'!\n", username)
	return nil
}

func handleRegisterUser(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <user_name>", cmd.name)
	}
	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return fmt.Errorf("unexpected error in GetUser: %w", err)
	}
	if err == nil {
		return fmt.Errorf("user '%s' is already registered", username)
	}
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC().UTC(),
		UpdatedAt: time.Now().UTC().UTC(),
		Name:      username,
	}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User '%s' successfully registered!\n", user.Name)
	printUser(user)
	return nil
}

func handleListUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error - handleUsers: users does not take arguments")
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error - handleUsers: %v", err)
	}
	for _, user := range users {
		if user.Name == *s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf("* ID:            %s\n", user.ID)
	fmt.Printf("* Created:       %v\n", user.CreatedAt)
	fmt.Printf("* Updated:       %v\n", user.UpdatedAt)
	fmt.Printf("* Name:          %s\n", user.Name)
}
