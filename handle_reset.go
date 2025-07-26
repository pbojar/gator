package main

import (
	"context"
	"fmt"
)

func handleReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error - handleReset: reset does not take arguments")
	}
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error - handleReset: %v", err)
	}
	return nil
}
