package api

import (
	"encoding/json"
	"net/http"

	"frontendmasters.com/go/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/")
	json.NewEncoder(w).Encode(data.GetAll())
}

