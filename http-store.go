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
	id := req.URL.Query()["id"][0]

	rows, err := db.Query("SELECT json FROM storage WHERE id = $1", id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{'status': 'nok', 'err': '%s'}", err)))
		return
	}
	defer rows.Close()
	var str string
	rows.Next()
	rows.Scan(&str)
	w.Write([]byte(str))
}

func DbPost(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	id := req.Form.Get("id")
	js, _ := ioutil.ReadAll(req.Body)
	_, err := db.Query("insert into storage values($1, $2) on conflict (id) do update set json = Excluded.json", id, js)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{'status': 'nok', 'err': '%s'}", err)))
		return
	}
	w.Write([]byte("{'status': 'ok'}"))

}

var db *sql.DB

func main() {
	port := os.Getenv("PORT")
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()
	http.HandleFunc("/", Index)
	http.HandleFunc("/dbget", DbGet)
	http.HandleFunc("/dbpost", DbPost)
	http.ListenAndServe(":"+port, nil)
}
