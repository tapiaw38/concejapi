package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/tapiaw38/concejapi/middlewares"
	"github.com/tapiaw38/concejapi/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers is the server
func HandlerServer() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
