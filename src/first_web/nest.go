package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")

	t.Execute(w, "Hello World!")
}
func main() {
	var err error
	if err != nil {
		panic(err)
	}
	server := http.Server{
		Addr: "127.0.0.1:9092",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
