package main

import (
	"encoding/json"
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName string
	ServerIP   string
}
type Serverslice struct {
	Servers []Server
}

func main() {
	//var s Serverslice
	/*
		str := `{"servers":[{"serverName":"Shanghai_VPN","ServerIP":"127.0.0.1"},
		{"serverName":"Beijing_VPN","ServerIP":"127.0.0.2"}]}`
	*/
	str1 := `{"Name":"Wednesday","Age":6,"Parents":["Gomez","Moriticia"]}`
	b := []byte(str1)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		panic(err)
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	//用simplejson解析JSON

	js, err := simplejson.NewJson(b)
	if err != nil {
		panic(err)
	}
	name, err := js.Get("Name").String()
	if err != nil {
		panic(err)
	}
	age, err := js.Get("Age").Int()
	if err != nil {
		panic(err)
	}
	arr, err := js.Get("Parents").Array()
	if err != nil {
		panic(err)
	}
	//slice := []string{}
	map1 := make(map[int]string)
	for key, pname := range arr {
		map1[key] = pname.(string)
		//slice = append(slice, pname.(string))
	}
	fmt.Println(map1)
	fmt.Println(age)
	fmt.Println(name)
	//fmt.Println(m)
}
