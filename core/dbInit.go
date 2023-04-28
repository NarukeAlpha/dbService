package core

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

/* dbConnection is a function that returns a sql.DB object dynamically, to be used in other functions
 */
func dbConnection() sql.DB {
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
	return *db
}

func SqlInit(wg *sync.WaitGroup) ([]DbMangaEntry, []DbChapterEntry) {
	defer wg.Done()
	var db sql.DB = dbConnection()
	defer db.Close()
	var mangaEntry []DbMangaEntry = readingMangaTable(db)

	var chapterEntry []DbChapterEntry = readingChapterListTable(db)

	//Loaded all mangas, now loading all chapters

	return mangaEntry, chapterEntry
}

func readingMangaTable(db sql.DB) []DbMangaEntry {
	var mangaL []DbMangaEntry
	query := "SELECT ID, Manga, LastChapter, Monitoring, Identifier FROM MangaList"
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
		var identifier string
		err := mangRows.Scan(&id, &manga, &lastChapter, &monitoring, &identifier)
		if err != nil {
			log.Printf("Coudln't scan row", err)
		}

		var entry = DbMangaEntry{
			Did:          id,
			Dmanga:       manga,
			DlastChapter: lastChapter,
			Dmonitoring:  monitoring,
			Didentifier:  identifier,
		}
		mangaL = append(mangaL, entry)
		log.Printf("id:%d ; manga: %s; lc: %d ; mon: %t \n", entry.Did, entry.Dmanga, entry.DlastChapter, entry.Dmonitoring, entry.Didentifier)
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
func updateMangaListTable(db sql.DB, entry DbMangaEntry) {

	var boolean int = 0
	/*
		turning off monitoring when manga is completed
	*/
	var query = fmt.Sprintf("UPDATE MangaList SET LastChapter = %d, Monitoring = %v WHERE ID = %d", entry.DlastChapter, boolean, entry.Did)
	_, err := db.ExecContext(context.Background(), query)
	if err != nil {
		log.Fatalf("failed to update manga list row:", err.Error())

	}
}
func addChapterListTable(db sql.DB, entry DbChapterEntry) {
	var boolean int = 0
	var query = fmt.Sprintf("INSERT INTO ChapterList (ID,Chapter,ChapterLink,released) VALUES (%v,%v,%d,%v)", entry.Did, entry.Dchapter, entry.DChapterLink, boolean)
	_, err := db.ExecContext(context.Background(), query)
	if err != nil {
		log.Fatalf("failed to add chapter list row:", err.Error())

	}
	boolean = 1
	var lc = entry.Dchapter - 1
	query = fmt.Sprintf("UPDATE ChapterList SET released = %v WHERE ID = %d", boolean, lc)
	_, err = db.ExecContext(context.Background(), query)
	if err != nil {
		log.Fatalf("failed to update old chapter list row:", err.Error())

	}
	query = fmt.Sprintf("UPDATE MangaList SET LastChapter = %d WHERE ID = %d", entry.Dchapter, entry.Did)
	_, err = db.ExecContext(context.Background(), query)
	if err != nil {
		log.Fatalf("failed to update manga list row:", err.Error())

	}

}
