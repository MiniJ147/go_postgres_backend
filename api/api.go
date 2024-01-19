package api

import "github.com/minij147/go_postgres_backend/internal/database"

type Config struct {
	DB *database.Queries
}

type Message struct {
	Msg string `json:"msg"`
}
