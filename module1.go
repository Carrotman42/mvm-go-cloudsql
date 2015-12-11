package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"google.golang.org/appengine"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root@unix(/cloudsql/project-id:region:instance-name)/"

var db = func() *sql.DB {
	r, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return r
}()

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := db.Ping(); err != nil {
		fmt.Fprintf(w, "error connecting with %s: %v", dsn, err)
		return
	}

	// Use db (but don't close it!)

	for _, q := range []string{
		"CREATE DATABASE IF NOT EXISTS x;",
		"CREATE TABLE IF NOT EXISTS x.x ( t1 int );",
		"INSERT INTO x.x VALUES (RAND()*100);",
	} {
		_, err := db.Exec(q)
		if err != nil {
			fmt.Fprintf(w, "error executing sql: %v", err)
			return
		}
	}

	r, err := db.Query("SELECT * from x.x;")
	if err != nil {
		fmt.Fprintf(w, "error executing sql: %v", err)
		return
	}
	defer r.Close()
	if cols, err := r.Columns(); err != nil {
		fmt.Fprintf(w, "couldn't read column names: %v<br>", err)
	} else {
		fmt.Fprintf(w, "%v<br>=====<br>", cols)
	}
	for r.Next() {
		var n int
		if err := r.Scan(&n); err != nil {
			fmt.Fprintf(w, "error scanning: %v<br>", err)
			break
		}
		fmt.Fprintf(w, "%v<br>", n)
	}
	if err := r.Err(); err != nil {
		fmt.Fprintf(w, "err in rows: %v<br>", err)
	}
}
