package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"frontendmasters.com/go/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


	id := r.URL.Query()["id"]
	if id != nil {
		index, err := strconv.Atoi(id[0])

		if err == nil && index < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[index])
		}else{
			http.Error(w, "Invalid Exibition", http.StatusBadRequest)
		}
	}else{
		json.NewEncoder(w).Encode(data.GetAll())

	}

}
