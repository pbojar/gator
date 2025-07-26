package main

import (
	"github.com/pbojar/gator/internal/config"
	"github.com/pbojar/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
