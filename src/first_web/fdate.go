package main

import (
	"html/template"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("fdate.html").Funcs(funcMap)
	t, _ = t.ParseFiles("fdate.html")
	t.Execute(w, time.Now())
}
func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:9091",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
