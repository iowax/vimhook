package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"
)

type result struct {
	Total   int           `json:"total"`
	Page    int           `json:"page"`
	Perpage int           `json:"per_page"`
	Paging  paging        `json:"paging"`
	Data    []interface{} `json:"data"`
}

type paging struct {
	Next     string  `json:"next"`
	Previous *paging `json:"previous"`
	First    string  `json:"first"`
	Last     string  `json:"last"`
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Search for videos related to: ")
	searchtext, _ := reader.ReadString('\n')

	searchURL := "https://api.vimeo.com/videos"

	req, _ := http.NewRequest("GET", searchURL, nil)

	q := req.URL.Query()
	q.Add("query", searchtext)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", "Bearer dc973e7482b29deee316fee415c2afa9")
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	data, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("The HTTP request failed with an error %s", err)
	} else {
		fmt.Println(string(data))
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	session1, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer session1.Close()

	session1.SetMode(mgo.Monotonic, true)

	c := session1.DB("vimeo").C("videos")
	var f result

	if c != nil {
		if err := json.Unmarshal(data, &f); err != nil {
			fmt.Println("JSON Parsing error")
		}
		if err := c.Insert(f.Paging); err != nil {
			fmt.Println("JSON Insertion error")
		}
	} else {
		fmt.Println("Couldn't connect to the database.")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}
