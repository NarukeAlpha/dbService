package core

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
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

func SqlInit() []DbMangaEntry {

	var db sql.DB = dbConnection()
	defer db.Close()
	var mangaEntry []DbMangaEntry = readingMangaTable(db)

	return mangaEntry
}

func readingMangaTable(db sql.DB) []DbMangaEntry {
	var mangaL []DbMangaEntry
	query := "SELECT ID, Manga, LastChapter, Monitoring, ChapterLink, Identifier FROM MasterTable"
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
		var chapterLink string
		var identifier string
		err := mangRows.Scan(&id, &manga, &lastChapter, &monitoring, &chapterLink, &identifier)
		if err != nil {
			log.Printf("Coudln't scan row", err)
		}

		var entry = DbMangaEntry{
			Did:          id,
			Dmanga:       manga,
			DlastChapter: lastChapter,
			Dmonitoring:  monitoring,
			DchapterLink: chapterLink,
			Didentifier:  identifier,
		}
		mangaL = append(mangaL, entry)
		log.Printf("id:%d ; manga: %s; lc: %d ; mon: %t ; chapLink: %s ; identifier: %s \n", entry.Did, entry.Dmanga, entry.DlastChapter, entry.Dmonitoring, entry.DchapterLink, entry.Didentifier)
	}
	if err := mangRows.Err(); err != nil {
		log.Fatal("Error iterating mangRows: ", err.Error())
	}
	return mangaL
}

func updateOffMangaListTable(db sql.DB, entry DbMangaEntry) {

	var boolean int = 0
	/*
		turning off monitoring when manga is completed
	*/
	var query = fmt.Sprintf("UPDATE MasterTable SET LastChapter = %d, Monitoring = %v WHERE ID = %d", entry.DlastChapter, boolean, entry.Did)
	_, err := db.ExecContext(context.Background(), query)
	if err != nil {
		log.Fatalf("failed to update manga list row:", err.Error())

	}
}
func addChapterToTable(db sql.DB, entry DbMangaEntry) {
	var query = fmt.Sprintf("UPDATE MasterTable SET LastChapter = %v WHERE ID = %v", entry.DlastChapter, entry.Did)
	_, err := db.ExecContext(context.Background(), query)
	if err != nil {
		log.Fatalf("failed to update Chapter List row:", err.Error())

	}

}
