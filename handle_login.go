package main

import "fmt"

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error - handleLogin: expected 1 argument")
	}
	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error - SetUser: %v", err)
	}
	fmt.Printf("Current user was set to '%s'!\n", username)
	return nil
}
