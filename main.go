package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()

	s := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("Server is running Port: ", port)

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
