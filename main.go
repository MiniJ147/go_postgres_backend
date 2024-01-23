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

	r.Route("/author", func(router chi.Router) {
		router.Get("/view", apiCfg.HandlerFetchAuthors)
		//router.Get("/find", apiCfg.HandlerFetchAuthorByName)
		router.Post("/create", apiCfg.HandlerCreateAuthor)
	})

	r.Route("/book", func(router chi.Router) {
		router.Post("/create", apiCfg.HandlerCreateBook)
		router.Get("/find", apiCfg.HandlerFetchBooks)
	})

	fmt.Println("Server Started!")

	http.ListenAndServe("localhost:3000", r)
}
