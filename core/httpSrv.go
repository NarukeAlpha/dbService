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

	return *rt
}
