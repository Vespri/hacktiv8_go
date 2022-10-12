package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Get Request
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	sb := string(body)
	log.Println(sb)

	// Post Request
	data := map[string]interface{}{
		"title":  "Kresna",
		"body":   "Vespri",
		"userID": 1,
	}

	requestJSON, err := json.Marshal(data)

	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	res1, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res1.Body.Close()

	body1, err := ioutil.ReadAll(res1.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body1))
}
