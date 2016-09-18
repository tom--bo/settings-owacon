package main

import (
	// "bytes"
	"database/sql"
	// "encoding/binary"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	ID          int
	AuthorID    int
	Keyword     string
	Description string
	UpdatedAt   time.Time
	CreatedAt   time.Time

	Html string
}

func pathURIEscape(s string) string {
	return (&url.URL{Path: s}).String()
}

func myhash(k string) string {
	return strings.Join(strings.Split(k, ""), "_") + ";"
}

func main() {
	//var r http.Request
	db, err := sql.Open("mysql", fmt.Sprintf(
		"isucon:isucon@tcp(localhost:3306)/isuda?loc=Local&parseTime=true",
	))
	if err != nil {
	}
	defer db.Close()

	rows, _ := db.Query(`
		SELECT id, author_id, keyword, description, updated_at, created_at FROM entry WHERE id < 7102 ORDER BY klen DESC, id DESC
	`)
	entries := make([]*Entry, 0, 500)

	// entry構造体の作成
	for rows.Next() {
		e := Entry{}
		rows.Scan(&e.ID, &e.AuthorID, &e.Keyword, &e.Description, &e.UpdatedAt, &e.CreatedAt)
		entries = append(entries, &e)
	}
	rows.Close()

	// keywordの配列の作成
	keywords := make([]string, 0, 500)
	for _, entry := range entries {
		keywords = append(keywords, (entry.Keyword))
	}

	for _, entry := range entries {
		kw2sha := make(map[string]string)
		for _, kw := range keywords {
			kw2sha[kw] = "isuda_" + fmt.Sprintf("%x", myhash((kw)))
			entry.Description = strings.Replace(entry.Description, kw, kw2sha[kw], -1)
		}
		entry.Description = html.EscapeString(entry.Description)
		for kw, hash := range kw2sha {
			u, _ := url.Parse("/keyword/" + pathURIEscape(kw))
			link := fmt.Sprintf("<a href=\"%s\">%s</a>", u, html.EscapeString(kw))
			entry.Description = strings.Replace(entry.Description, hash, link, -1)
		}
		entry.Description = strings.Replace(entry.Description, "\n", "<br />\n", -1)

		// update idでdescription
		db.Exec(
			"UPDATE entry SET description = ? WHERE id = ?",
			//"INSERT INTO entry(klen) values(?) WHERE id = ?",
			entry.Description, entry.ID,
		)
	}
}

//db, err := sql.Open("mysql", "root:root@unix(/var/run/mysql/mysql.sock)/isucon")

//	db, err = sql.Open("mysql", fmt.Sprintf(
//		"isucon:isucon@tcp(localhost:3306)/isuda?loc=Local&parseTime=true"
//	))
//	if err != nil {panic(err.Error())}
//	defer db.Close()
//
//	rows, err2 := db.Query(`
//		SELECT * FROM entry ORDER BY CHARACTER_LENGTH(keyword) DESC
//	`)
//	if err2 != nil {panic(err2.Error())}
//
//	entries := make([]*Entry, 0, 500)
//	for rows.Next() {
//		e := Entry{}
//		err3 := rows.Scan(&e.ID, &e.AuthorID, &e.Keyword, &e.Description, &e.UpdatedAt, &e.CreatedAt)
//		if err3 != nil {panic(err3.Error())}
//		entries = append(entries, &e)
//	}
//	rows.Close(
//
//
//	keywords := make([]string, 0, 500)
//	for _, entry := range entries {
//		keywords = append(keywords, regexp.QuoteMeta(entry.Keyword))
//	}
//	re := regexp.MustCompile("(" + strings.Join(keywords, "|") + ")")
//	kw2sha := make(map[string]string)
//	content = re.ReplaceAllStringFunc(content, func(kw string) string {
//		kw2sha[kw] = "isuda_" + fmt.Sprintf("%x", sha1.Sum([]byte(kw)))
//		return kw2sha[kw]
//	})
//	content = html.EscapeString(content)
//	for kw, hash := range kw2sha {
//		u, err4 := r.URL.Parse(baseUrl.String() + "/keyword/" + pathURIEscape(kw))
//		if err4 != nil {panic(err4.Error())}
//		link := fmt.Sprintf("<a href=\"%s\">%s</a>", u, html.EscapeString(kw))
//		content = strings.Replace(content, hash, link, -1)
//	}
//	return strings.Replace(content, "\n", "<br />\n", -1)

//	rows, err := db.Query("SELECT id,content FROM memos")
//	defer rows.Close()
//	if err != nil {
//		panic(err.Error())
//	}
//
//	for rows.Next() {
//		var id int64
//		var content string
//
//		if err := rows.Scan(&id, &content); err != nil {
//			panic(err.Error())
//		}
//		title := strings.SplitN(content, "\n", 2)[0]
//
//		_, err := db.Exec(
//			"UPDATE memos SET title = ? WHERE id = ?",
//			title, id,
//		)
//		if err != nil {
//			panic(err.Error())
//		}
//	}

func main1() {
	db, err := sql.Open("mysql", "root:root@unix(/var/lib/mysql/mysql.sock)/isucon")
	if err != nil {
		panic(err.Error())

	}
	defer db.Close()

	rows, err := db.Query("SELECT id,content FROM memos")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int64
		var content string

		if err := rows.Scan(&id, &content); err != nil {
			panic(err.Error())
		}
		title := strings.SplitN(content, "\n", 2)[0]

		_, err := db.Exec(
			"UPDATE memos SET title = ? WHERE id = ?",
			title, id,
		)
		if err != nil {
			panic(err.Error())
		}
	}

}

func main2() {
	db, err := sql.Open("mysql", "root@/isuconp")
	if err != nil {
		panic(err.Error())

	}
	defer db.Close()

	rows, err := db.Query("SELECT id,mime,imgdata FROM posts")
	defer rows.Close()
	if err != nil {
		panic(err.Error())

	}

	for rows.Next() {
		var id int64
		var mime string
		var imgdata []byte
		if err := rows.Scan(&id, &mime, &imgdata); err != nil {
			panic(err.Error())
		}
		mime2 := strings.Split(mime, "/")
		filename := strconv.FormatInt(id, 10) + "." + mime2[1]

		file, err2 := os.Create(filename)
		if err2 != nil {
			fmt.Println("file create err:", err2)
			return
		}

		_, err3 := file.Write(imgdata)
		if err3 != nil {
			fmt.Println("file write err:", err3)
			return
		}

	}

}
