package main

import (
	// "bytes"
	"database/sql"
	// "encoding/binary"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"strings"
)

func main() {
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
