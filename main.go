package main

import (
	"dBService/core"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"net/http"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	var MangaList, ChapterList = core.SqlInit(&wg)
	wg.Wait()

	log.Println(MangaList, ChapterList)

	r := core.InitHttpServerMux(MangaList, ChapterList)

	//s := http.NewServeMux()
	//s.Handler =
	//
	//	http.HandleFunc("/mangas", func(w http.ResponseWriter, r *http.Request) {
	//		if err := json.NewEncoder(w).Encode(MangaList); err != nil {
	//			http.Error(w, err.Error(), http.StatusInternalServerError)
	//			return
	//		}
	//	})
	//http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
	//	os.Exit(2)
	//
	//})
	//
	http.ListenAndServe("localhost:8080", &r)

}
