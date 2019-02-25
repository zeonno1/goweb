package main

import (
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Co",
		HttpOnly: true,
	}
	w.Header().Set("Set-cookie", c1.String())
	w.Header().Set("Set-cookie", c2.String())
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:9091",
	}
	http.HandleFunc("/set_cookie", setCookie)
	server.ListenAndServe()
}
