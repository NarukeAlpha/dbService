package main

import (
	"dBService/core"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"net/http"
)

func main() {

	var MangaList = core.SqlInit()

	log.Println(MangaList)

	r := core.InitHttpServerMux(MangaList)

	http.ListenAndServe("localhost:8080", &r)

}
