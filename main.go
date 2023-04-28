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

	http.ListenAndServe("localhost:8080", &r)

}
