package core

type DbMangaEntry struct {
	Did          int    `json:"did"`
	Dmanga       string `json:"dmanga"`
	DlastChapter int    `json:"dlastChapter"`
	Dmonitoring  bool   `json:"dmonitoring"`
	Didentifier  string `json:"didentifier"`
}

type DbChapterEntry struct {
	Did          int    `json:"did"`
	Dchapter     int    `json:"dchapter"`
	DChapterLink string `json:"d_chapter_link"`
	Dreleased    bool   `json:"dreleased"`
}
