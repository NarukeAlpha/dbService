package core

type DbMangaEntry struct {
	Did          int
	Dmanga       string
	DlastChapter int
	Dmonitoring  bool
}

type DbChapterEntry struct {
	Did          int
	Dchapter     int
	DChapterLink string
	Dreleased    bool
}
