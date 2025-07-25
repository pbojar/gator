package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pbojar/gator/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("error parsing args:\nusage: ./gator <command> (args...)\n")
		os.Exit(1)
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
	}

	s := state{&cfg}
	c := commands{
		commands: make(map[string]func(*state, command) error),
	}
	c.register("login", handleLogin)
	err = c.run(&s, cmd)
	if err != nil {
		fmt.Printf("error running command: %v\n", err)
		os.Exit(1)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
	}

	cfgMarshalled, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Printf("Config contents:\n%s\n", string(cfgMarshalled))
}
