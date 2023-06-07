package main

import (
	"dBService/core"
	_ "github.com/denisenkom/go-mssqldb"
	"net/http"
)

func main() {

	var MangaList = core.SqlInit()

	r := core.InitHttpServerMux(MangaList)

	http.ListenAndServe("localhost:8080", &r)

}
