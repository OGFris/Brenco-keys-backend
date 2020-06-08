package handler

import (
	"encoding/json"
	"github.com/OGFris/Brenco-keys-backend/database"
	"log"
	"net/http"
)

func KeysGet(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys, err := db.GetKeys()
		if err != nil {
			log.Printf("error from %s: %v", r.RemoteAddr, err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		bytes, err := json.Marshal(keys)
		if err != nil {
			log.Printf("error from %s: %v", r.RemoteAddr, err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.Write(bytes)
		w.WriteHeader(http.StatusOK)
	}
}
