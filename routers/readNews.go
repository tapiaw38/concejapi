package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tapiaw38/concejapi/db"
)

func ReadNews(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The news id is invalid", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("paginate")) > 1 {
		http.Error(w, "You must send the parameter for paging", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("paginate"))
	if err != nil {
		http.Error(w, "You must send the pagination parameter with a value greater than 0", http.StatusBadRequest)
		return
	}

	pageInt := int64(page)

	response, correct := db.ReadNews(ID, pageInt)

	if correct == false {
		http.Error(w, "Error reading the news", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
