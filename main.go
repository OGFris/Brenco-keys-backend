package main

import (
	"fmt"
	"github.com/OGFris/Brenco-keys-backend/database"
	"github.com/OGFris/Brenco-keys-backend/handler"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := database.New()
	if err != nil {

		panic(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/keys/create", handler.CreatePost(db)).Methods("POST")
	router.HandleFunc("/api/keys/remove/{id}", handler.RemovePost(db)).Methods("POST")
	router.HandleFunc("/api/keys", handler.KeysGet(db)).Methods("GET")

	s := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("Server is running Port: ", port)

	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
