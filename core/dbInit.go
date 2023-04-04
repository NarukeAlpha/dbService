package core

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func SqlInit() ([]DbMangaEntry, []DbChapterEntry) {

	//loading sql key from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env file")
	}
	connString := fmt.Sprintf(os.Getenv("dbkey"))

	//sql server connection
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()
	var mangaEntry []DbMangaEntry = readingMangaTable(*db)

	var chapterEntry []DbChapterEntry = readingChapterListTable(*db)

	//Loaded all mangas, now loading all chapters

	return mangaEntry, chapterEntry
}

func readingMangaTable(db sql.DB) []DbMangaEntry {
	var mangaL []DbMangaEntry
	query := "SELECT ID, Manga, LastChapeter, Monitoring FROM MangaList"
	mangRows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal("Error querying database: ", err.Error())
	}
	defer mangRows.Close()

	for mangRows.Next() {

		var id int
		var manga string
		var lastChapter int
		var monitoring bool
		err := mangRows.Scan(&id, &manga, &lastChapter, &monitoring)
		if err != nil {
			log.Printf("Coudln't scan row", err)
		}

		var entry = DbMangaEntry{
			Did:          id,
			Dmanga:       manga,
			DlastChapter: lastChapter,
			Dmonitoring:  monitoring,
		}
		mangaL = append(mangaL, entry)
		log.Printf("id:%d ; manga: %s; lc: %d ; mon: %t \n", entry.Did, entry.Dmanga, entry.DlastChapter, entry.Dmonitoring)
	}
	if err := mangRows.Err(); err != nil {
		log.Fatal("Error iterating mangRows: ", err.Error())
	}
	return mangaL
}

func readingChapterListTable(db sql.DB) []DbChapterEntry {
	var chapterL []DbChapterEntry
	query := "SELECT ID,Chapter,ChapterLink,Released FROM ChapterList"
	chRows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatalf("failed getting chapter list row:", err.Error())
	}
	defer chRows.Close()

	for chRows.Next() {
		var id int
		var chapter int
		var chapterlink string
		var released bool
		err := chRows.Scan(&id, &chapter, &chapterlink, &released)
		if err != nil {
			log.Printf("failed to scan over row : ", err)
		}
		var entry = DbChapterEntry{
			Did:          id,
			Dchapter:     chapter,
			DChapterLink: chapterlink,
			Dreleased:    released,
		}

		if entry.Dreleased == false {
			chapterL = append(chapterL, entry)
			log.Printf("id:%d ; chapter: %d; cl: %s ; mon: %t \n", entry.Did, entry.Dchapter, entry.DChapterLink, entry.Dreleased)
		}
	}
	return chapterL
}
