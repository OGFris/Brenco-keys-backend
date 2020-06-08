package main

import (
	"fmt"
	"github.com/OGFris/Brenco-keys-backend/database"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/pgtype"
	_ "github.com/jackc/pgx/stdlib"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	_, err := database.New()
	if err != nil {

		panic(err)
	}

	router := mux.NewRouter()

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
