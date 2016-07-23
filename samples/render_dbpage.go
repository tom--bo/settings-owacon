package main

import (
	// "bytes"
	"database/sql"
	// "encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
	"github.com/russross/blackfriday"
)

func main() {
        fmap       := template.FuncMap{
                "url_for": func(path string) string {
                        return "/" + path
                },
                "first_line": func(s string) string {
                        sl := strings.Split(s, "\n")
                        return sl[0]
                },
                "gen_markdown": func(s string) template.HTML {
			output:=blackfriday.MarkdownBasic([]byte(s))
                        return template.HTML(output)
                },
        }


	tmpl := template.Must(template.New("tmpl").Funcs(fmap).ParseGlob("templates/*.html"))
os.Exit(0);

	db, err := sql.Open("mysql", "isucon@/isucon")
	if err != nil {
		panic(err.Error())

	}
	defer db.Close()
	rows, err := db.Query("SELECT id, user, content, is_private, created_at, updated_at FROM memos")
	defer rows.Close()
	if err != nil {
		panic(err.Error())

	}

	for rows.Next() {
		var id int64
		var mime string
		var imgdata []byte
		if err := rows.Scan(&memo.Id, &memo.User, &memo.Content, &memo.IsPrivate, &memo.CreatedAt, &memo.UpdatedAt); err != nil {
			panic(err.Error())
		}
		mime2 := strings.Split(mime, "/")
		filename := strconv.FormatInt(id, 10) + "." + mime2[1]

		file, err2 := os.Create(filename)
		if err2 != nil {
			fmt.Println("file create err:", err2)
			return
		}

		v := &View{
			User:    user,
			Memo:    memo,
			Older:   older,
			Newer:   newer,
			Session: session,
		}

		   if err = tmpl.ExecuteTemplate(file, "memo", v); err != nil {
			   serverError(w, err)
		   }

	}
}
