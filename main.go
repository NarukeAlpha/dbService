package main

import (
	"dBService/core"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func main() {
	var MangaList, ChapterList = core.SqlInit()

	log.Println(MangaList, ChapterList)
}
