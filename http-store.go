package main

import (
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "view/index.html")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/"+port, Index)
	http.ListenAndServe(":80", nil)
}
