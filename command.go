package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdFunc, exists := c.commands[cmd.name]
	if !exists {
		return fmt.Errorf("error - run: command '%s' is not supported", cmd.name)
	}
	return cmdFunc(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
