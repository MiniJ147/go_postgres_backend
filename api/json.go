package api

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
