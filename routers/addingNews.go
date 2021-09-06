package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tapiaw38/concejapi/db"
	"github.com/tapiaw38/concejapi/models"
)

func AddingNews(w http.ResponseWriter, r *http.Request) {

	var n models.AddNews

	err := json.NewDecoder(r.Body).Decode(&n)

	NewRegister := models.News{
		User:     UserID,
		Title:    n.Title,
		Body:     n.Body,
		Category: n.Category,
		Picture:  n.Picture,
		Date:     time.Now(),
	}

	_, status, err := db.InsertNews(NewRegister)
	if err != nil {
		http.Error(w, "An error occurred when trying to enter a news item "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "It was not possible to add a news ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
