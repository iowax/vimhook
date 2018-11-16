package main

// func main() {
// 	fmt.Print("Search for videos related to: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	searchtext, _ := reader.ReadString('\n')

// 	client, err := vimeo.NewVimeoClient()
// }

// type result struct {
// 	Total   int     `json:"total"`
// 	Page    int     `json:"page"`
// 	Perpage int     `json:"per_page"`
// 	Paging  paging  `json:"paging"`
// 	Data    []video `json:"data"`
// }

// type paging struct {
// 	Next     string  `json:"next"`
// 	Previous *paging `json:"previous"`
// 	First    string  `json:"first"`
// 	Last     string  `json:"last"`
// }

// type video struct {
// 	URI string `json:"uri"`
// }

// func main() {

// 	//reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Search for videos related to: ")
// 	//searchtext, _ := reader.ReadString('\n')

// 	searchURL := "https://api.vimeo.com/videos?query=cats&fields=uri"

// 	req, _ := http.NewRequest("GET", searchURL, nil)

// 	q := req.URL.Query()
// 	//q.Add("query", searchtext)
// 	req.URL.RawQuery = q.Encode()

// 	req.Header.Set("Authorization", "Bearer dc973e7482b29deee316fee415c2afa9")
// 	req.Header.Set("Content-Type", "application/json")

// 	client := http.Client{}
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	data, _ := ioutil.ReadAll(resp.Body)

// 	defer resp.Body.Close()

// 	if err != nil {
// 		fmt.Printf("The HTTP request failed with an error %s", err)
// 	} else {
// 		fmt.Println(string(data))
// 	}

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	session1, err := mgo.Dial("mongodb://localhost:27017")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	defer session1.Close()

// 	session1.SetMode(mgo.Monotonic, true)

// 	c := session1.DB("vimeo").C("uris")
// 	var f result

// 	if c != nil {
// 		if err := json.Unmarshal(data, &f); err != nil {
// 			fmt.Println("JSON Parsing error")
// 		}
// 		for _, val := range f.Data {
// 			fmt.Println(val)
// 		}
// 		// if err := c.Insert(val); err != nil {
// 		// 	fmt.Println("JSON Insertion error")
// 		// }
// 	} else {
// 		fmt.Println("Couldn't connect to the database.")
// 	}

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }
