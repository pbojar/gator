package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pbojar/gator/internal/database"
)

func handleRegister(s *state, cmd command) error {
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}
	_, err = s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error - database.CreateUser: %v", err)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error - SetUser: %v", err)
	}
	fmt.Printf("User '%s' registered!\n", username)
	fmt.Printf("\tuuid - %s\n", params.ID.String())
	fmt.Printf("\tt_cr - %s\n", params.CreatedAt.String())
	fmt.Printf("\tt_up - %s\n", params.UpdatedAt.String())
	fmt.Printf("\tname - %s\n", params.Name)
	return nil
}
