package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/pbojar/gator/internal/config"
	"github.com/pbojar/gator/internal/database"
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
	db, err := sql.Open("postgres", *cfg.DBURL)
	if err != nil {
		fmt.Printf("error opening db: %v\n", err)
	}
	dbQueries := database.New(db)

	s := state{
		db:  dbQueries,
		cfg: &cfg,
	}
	c := commands{
		commands: make(map[string]func(*state, command) error),
	}
	c.register("login", handleLogin)
	c.register("register", handleRegisterUser)
	c.register("reset", handleReset)
	c.register("users", handleListUsers)
	c.register("agg", handleAgg)
	c.register("addfeed", loggedIn(handleAddfeed))
	c.register("feeds", handleListFeeds)
	c.register("follow", loggedIn(handleFollow))
	c.register("following", handleFollowing)
	c.register("unfollow", loggedIn(handleUnfollow))
	err = c.run(&s, cmd)
	if err != nil {
		fmt.Printf("error running command: %v\n", err)
		os.Exit(1)
	}
}
