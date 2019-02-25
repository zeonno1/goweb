package main

import (
	"fmt"
)

func main() {
	b := []byte("abc")
	fmt.Println(b)
	s := string(b)
	fmt.Println(s)
}
