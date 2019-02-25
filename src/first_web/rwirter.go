package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web Programming</titel></head>
	<body><h1>Hello World!</h1></body>
	</html>`
	w.Write([]byte(str))
}
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service try net door")
}
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	post := &Post{
		User:    "sam",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:9092",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/header", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
