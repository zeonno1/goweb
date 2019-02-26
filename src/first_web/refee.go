package main

import (
	"fmt"
	"time"
)

func dateFormat(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}
func main() {
	b := []byte("abc")
	fmt.Println(b)
	s := string(b)
	fmt.Println(s)
	fmt.Println(time.Now().Unix())
	fmt.Println(dateFormat(time.Now()))
}
