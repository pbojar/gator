package main

import (
	"context"
	"fmt"
)

func handleUsers(s *state, cmd command) error {
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
