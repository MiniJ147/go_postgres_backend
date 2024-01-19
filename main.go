package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/minij147/go_postgres_backend/api"
	"github.com/minij147/go_postgres_backend/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		fmt.Print("ERROR GRABBING DB URL")
	}
	//sql

	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	queries := database.New(conn)

	apiCfg := api.Config{
		DB: queries,
	}

	//routing
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		api.SendJSON(w, r, 200, api.Message{Msg: "Hello From Server!"})
	})

	r.Post("/create", apiCfg.HandlerCreateAuthor)

	fmt.Println("Server Started!")

	http.ListenAndServe("localhost:3000", r)
}
