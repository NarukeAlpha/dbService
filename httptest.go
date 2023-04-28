package main

import (
	"bytes"
	"net/http"
)

func main() {

	//type DbMangaEntry struct {
	//	Did          int    `json:"did"`
	//	Dmanga       string `json:"dmanga"`
	//	DlastChapter int    `json:"dlastChapter"`
	//	Dmonitoring  bool   `json:"dmonitoring"`
	//	Didentifier  string `json:"didentifier"`
	//}
	//
	//var did int = 1
	//var dmanga string = "test"
	//var dlastChapter int = 1
	//var dmonitoring bool = true
	//var didentifier string = "test"
	//
	//var test = DbMangaEntry{
	//	Did:          did,
	//	Dmanga:       dmanga,
	//	DlastChapter: dlastChapter,
	//	Dmonitoring:  dmonitoring,
	//	Didentifier:  didentifier,
	//}

	/*payload := []byte(`{"Did": 55, "Dmanga": "MangaUrl", "DlastChapter": 200, "Dmonitoring": true, "Didentifier": "MangaChapterDate"}`)
	url := "http://localhost:8080/update-MangaList"
	*/
	payload := []byte(`{"Did": 1, "Dchapter": "241", "DChapterLink": "https://readeleceed.com/manga/eleceed-chapter-241/", "Dreleased": false}`)
	url := "http://localhost:8080/update-ChapterList"

	resp, err := sendPostRequest(url, payload)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// do something with the response body, if necessary

}

func sendPostRequest(url string, jsonPayload []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
