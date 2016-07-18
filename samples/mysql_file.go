package main

import (
	// "bytes"
	"database/sql"
	// "encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
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
