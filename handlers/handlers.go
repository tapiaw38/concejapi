package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tapiaw38/concejapi/middlewares"
	"github.com/tapiaw38/concejapi/routers"
)

// HandlerServer is the server
func HandlerServer() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateToken(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateToken(routers.ModifiedProfile))).Methods("PUT")
	router.HandleFunc("/news", middlewares.CheckDB(middlewares.ValidateToken(routers.AddingNews))).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
