package main

import "net/http"

func Index(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "view/index.html")
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":80", nil)
}
