package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}
type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "ShangHai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "BeiJing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	var f interface{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
