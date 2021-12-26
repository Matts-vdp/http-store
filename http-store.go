package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func Index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hallo"))
}

func DbGet(w http.ResponseWriter, req *http.Request) {
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS storage (id varchar(20) primary key, json text)"); err != nil {
		fmt.Printf("Error creating database table: %s", err)
		return
	}
	id := req.URL.Query()["id"][0]

	rows, err := db.Query("SELECT json FROM storage WHERE id = " + id)
	if err != nil {
		fmt.Println("cant get", id)
		return
	}
	defer rows.Close()
	var str string
	rows.Scan(&str)
	w.Write([]byte(str))
}

func DbPost(w http.ResponseWriter, req *http.Request) {
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS storage (id varchar(20) primary key, json text)"); err != nil {
		fmt.Printf("Error creating database table: %q", err)
		return
	}
	id := req.URL.Query()["id"][0]
	js, _ := ioutil.ReadAll(req.Body)
	q := fmt.Sprintf("INSERT INTO storage VALUES (%s, %s)", id, js)
	if _, err := db.Exec(q); err != nil {
		fmt.Printf("Error inserting: %s", id)
		return
	}
	w.Write([]byte("{'status', 'ok'}"))
}

var db *sql.DB

func main() {
	port := os.Getenv("PORT")
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	http.HandleFunc("/", Index)
	http.HandleFunc("/dbget", DbGet)
	http.HandleFunc("/dbpost", DbPost)
	http.ListenAndServe(":"+port, nil)
}
