package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/minij147/go_postgres_backend/internal/database"
)

func (apiCfg *Config) HandlerCreateBook(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Title    string        `json:"title"`
		AuthorID uuid.NullUUID `json:"authorID"`
	}

	param := Params{}
	json.NewDecoder(r.Body).Decode(&param)

	fmt.Println(param.Title, param.AuthorID)

	book, err := apiCfg.DB.CreateBook(r.Context(), database.CreateBookParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     param.Title,
		AuthorID:  param.AuthorID,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Could not create user")
		return
	}

	SendJSON(w, r, http.StatusAccepted, book)
}

func (apiCfg *Config) HandlerFetchBooks(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Books []database.Book `json:"books"`
	}

	fmt.Println(r.URL.Query().Get("name"))

	queryResult := r.URL.Query().Get("name")
	if queryResult == "" {
		//return all books
		books, err := apiCfg.DB.FetchBooks(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println("COuld not find all books")
		}

		SendJSON(w, r, http.StatusAccepted, Response{
			Books: books,
		})

		return
	}
	//else we don't have a query
	requestAuthor, err := apiCfg.DB.FetchAuthorByName(r.Context(), queryResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("COuld not find AUthor")
	}

	books, err := apiCfg.DB.FetchBooksByAuthorName(r.Context(), uuid.NullUUID{
		UUID:  requestAuthor.ID,
		Valid: true,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("COuld not find books")
	}

	SendJSON(w, r, http.StatusAccepted, Response{
		Books: books,
	})
}
