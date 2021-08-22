package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/concejapi/models"
)

func AddingNews(w http.ResponseWriter, r *http.Request) {

	var n models.News

	err := json.NewDecoder(n.Body).Decode(&n)
}
