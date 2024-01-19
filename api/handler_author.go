package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/minij147/go_postgres_backend/internal/database"
)

func (apiCfg *Config) HandlerCreateAuthor(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Name string `json:"name"`
	}

	params := Params{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if params.Name == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := apiCfg.DB.CreateAuthor(r.Context(), database.CreateAuthorParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Could not create user")
		return
	}

	SendJSON(w, r, http.StatusAccepted, user)
}
