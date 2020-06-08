package handler

import (
	"github.com/OGFris/Brenco-keys-backend/database"
	"log"
	"net/http"
)

func CreatePost(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")

		if name == "" {
			log.Printf("bad create post from %s: empty name", r.RemoteAddr)
			w.Write([]byte("bad create post"))
			w.WriteHeader(http.StatusBadRequest)

			return
		}

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
