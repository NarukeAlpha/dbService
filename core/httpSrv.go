package core

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func InitHttpServerMux(mL []DbMangaEntry, cL []DbChapterEntry) mux.Router {
	rt := mux.NewRouter()
	rt.HandleFunc("/get-MangaList", func(w http.ResponseWriter, request *http.Request) {
		if err := json.NewEncoder(w).Encode(mL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	rt.HandleFunc("/get-ChapterList", func(w http.ResponseWriter, request *http.Request) {
		if err := json.NewEncoder(w).Encode(mL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	rt.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(2)
	})

	rt.HandleFunc("/update-MangaList", func(w http.ResponseWriter, r *http.Request) {
		var mangaEntry DbMangaEntry
		if err := json.NewDecoder(r.Body).Decode(&mangaEntry); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mL = append(mL, mangaEntry)
		var db = dbConnection()
		defer db.Close()
		updateMangaListTable(db, mangaEntry)

	}).Methods("POST")

	return *rt
}
