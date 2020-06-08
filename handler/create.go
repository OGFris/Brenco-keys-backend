package handler

import (
	"github.com/OGFris/Brenco-keys-backend/database"
	"log"
	"net/http"
)

func CreatePost(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Printf("bad create post from %s: %v", r.RemoteAddr, err)
			w.Write([]byte("bad create post"))
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		name := r.FormValue("name")

		key, err := db.CreateKey(name)
		if err != nil {
			log.Printf("error from %s: %v", r.RemoteAddr, err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.Write([]byte(key))
		w.WriteHeader(http.StatusOK)
	}
}
