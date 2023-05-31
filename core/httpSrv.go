package core

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func InitHttpServerMux(mL []DbMangaEntry) mux.Router {
	rt := mux.NewRouter()
	rt.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(2)
	})
	rt.HandleFunc("/MangaList", func(w http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			if err := json.NewEncoder(w).Encode(mL); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "PUT":
			var mangaEntry DbMangaEntry
			if err := json.NewDecoder(request.Body).Decode(&mangaEntry); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			mL = append(mL, mangaEntry)
			var db = dbConnection()
			defer db.Close()
			addChapterToTable(db, mangaEntry)

		case "POST":
			var mangaEntry DbMangaEntry
			if err := json.NewDecoder(request.Body).Decode(&mangaEntry); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			mL = append(mL, mangaEntry)
			var db = dbConnection()
			defer db.Close()
			addNewMangaToTable(db, mangaEntry)
		}
	}).Methods("GET", "PUT", "POST")

	return *rt
}
