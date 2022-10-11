package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "https://developer.com:80/hi?name=kresna&age=21"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("URL : %s\n", urlString)

	fmt.Printf("Protocol : %s\n", u.Scheme)
	fmt.Printf("Host : %s\n", u.Host)
	fmt.Printf("Path : %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("Name: %s, Age: %s\n", name, age)
}
