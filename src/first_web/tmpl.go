package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t1, _ := template.ParseFiles("tmpl.html")
	dayOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t1.Execute(w, dayOfWeek)
}
func main() {
	var err error
	if err != nil {
		panic(err)
	}
	server := http.Server{
		Addr: "127.0.0.1:9091",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
