package main

//func addChapterToTable(db sql.DB, entry DbChapterEntry) {
//	var query = fmt.Sprintf("INSERT INTO ChapterList (ID,Chapter,ChapterLink,released) VALUES (%v,%v,'%v',%v)", entry.Did, entry.Dchapter, entry.DChapterLink, boolean)
//	_, err := db.ExecContext(context.Background(), query)
//	if err != nil {
//		log.Fatalf("failed to insert chapter list row:", err.Error())
//
//	}
//	boolean = 1
//	var lc = entry.Dchapter - 1
//	query = fmt.Sprintf("UPDATE ChapterList SET released = %v WHERE ID = %v AND Chapter = %v", boolean, entry.Did, lc)
//	_, err = db.ExecContext(context.Background(), query)
//	if err != nil {
//		log.Fatalf("failed to update old chapter list row:", err.Error())
//
//	}
//}

/*reading chapter list no longer used after merged tables into a master list.  Keeping for reference
*
 */

//	func readingChapterListTable(db sql.DB) []DbChapterEntry {
//		var chapterL []DbChapterEntry
//		query := "SELECT ID,Chapter,ChapterLink,Released FROM ChapterList"
//		chRows, err := db.QueryContext(context.Background(), query)
//		if err != nil {
//			log.Fatalf("failed getting chapter list row:", err.Error())
//		}
//		defer chRows.Close()
//
//		for chRows.Next() {
//			var id int
//			var chapter int
//			var chapterlink string
//			var released bool
//			err := chRows.Scan(&id, &chapter, &chapterlink, &released)
//			if err != nil {
//				log.Printf("failed to scan over row : ", err)
//			}
//			var entry = DbChapterEntry{
//				Did:          id,
//				Dchapter:     chapter,
//				DChapterLink: chapterlink,
//				Dreleased:    released,
//			}
//
//			if entry.Dreleased == false {
//				chapterL = append(chapterL, entry)
//				log.Printf("id:%d ; chapter: %d; cl: %s ; mon: %t \n", entry.Did, entry.Dchapter, entry.DChapterLink, entry.Dreleased)
//			}
//		}
//		return chapterL
//	}

//rt.HandleFunc("/ChapterList", func(w http.ResponseWriter, request *http.Request) {
//		switch request.Method {
//		case "GET":
//			if err := json.NewEncoder(w).Encode(mL); err != nil {
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//		case "PUT":
//			var chapterEntry DbChapterEntry
//			if err := json.NewDecoder(request.Body).Decode(&chapterEntry); err != nil {
//				http.Error(w, err.Error(), http.StatusBadRequest)
//				return
//			}
//			cL = append(cL, chapterEntry)
//			var db = dbConnection()
//			defer db.Close()
//			addChapterToTable(db, chapterEntry)
//		case "POST":
//			fmt.Println("POST still in dev")
//
//		}
//	}).Methods("GET", "PUT", "POST")
