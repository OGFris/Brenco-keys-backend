package handler

import (
	"github.com/OGFris/Brenco-keys-backend/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func RemovePost(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			log.Printf("bad remove post from %s: %v", r.RemoteAddr, err)
			w.Write([]byte("bad remove post"))
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		err = db.RemoveKey(id)
		if err != nil {
			log.Printf("error from %s: %v", r.RemoteAddr, err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
