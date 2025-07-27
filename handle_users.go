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
		return fmt.Errorf("error - handleLogin: expected 1 argument")
	}
	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return fmt.Errorf("error - handleLogin: '%s' is not a registered user", username)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error - SetUser: %v", err)
	}
	fmt.Printf("Current user was set to '%s'!\n", username)
	return nil
}

func handleRegisterUser(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error - handleRegister: expected 1 argument")
	}
	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return fmt.Errorf("error - database.GetUser: %v", err)
	}
	if err == nil {
		return fmt.Errorf("error - handleRegister: user '%s' is already registered", username)
	}
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC().UTC(),
		UpdatedAt: time.Now().UTC().UTC(),
		Name:      username,
	}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error - database.CreateUser: %v", err)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error - SetUser: %v", err)
	}
	fmt.Printf("User '%s' registered!\n", username)
	fmt.Printf("\tuuid - %s\n", user.ID.String())
	fmt.Printf("\tt_cr - %s\n", user.CreatedAt.String())
	fmt.Printf("\tt_up - %s\n", user.UpdatedAt.String())
	fmt.Printf("\tname - %s\n", user.Name)
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
