package main

import (
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hallo"))
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":"+port, nil)
}
