package main

import (
	"context"
	"fmt"
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
