package vimeo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

//Client struct to handle communication with vimeo api
type Client struct {
	client *http.Client
	dbconn *DBConnector
}

type VideoResponse struct {
	Embed struct {
		HTML string `json:"html"`
	} `json:"embed"`
}
type VimeoVideoConfig struct {
	Request struct {
		Files struct {
			Hls struct {
				SeparateAv bool   `json:"separate_av"`
				DefaultCdn string `json:"default_cdn"`
				Cdns       struct {
					Level3 struct {
						URL    string `json:"url"`
						Origin string `json:"origin"`
					} `json:"level3"`
				} `json:"cdns"`
				Origin string `json:"origin"`
			} `json:"hls"`
			Progressive []struct {
				Profile int    `json:"profile"`
				Width   int    `json:"width"`
				Mime    string `json:"mime"`
				Fps     int    `json:"fps"`
				URL     string `json:"url"`
				Cdn     string `json:"cdn"`
				Quality string `json:"quality"`
				ID      int    `json:"id"`
				Origin  string `json:"origin"`
				Height  int    `json:"height"`
			} `json:"progressive"`
		} `json:"files"`
		Lang    string `json:"lang"`
		Country string `json:"country"`
	} `json:"request"`
}

//NewVimeoClient provides an instance of a VimeoClient to connect to vimeo API
func NewVimeoClient() (*Client, error) {
	restclient := http.DefaultClient
	dbinstance, err := NewDBConnector(defaultMgoURL)

	if err != nil {
		return nil, err
	}
	return &Client{client: restclient, dbconn: dbinstance}, nil

}

func (c *Client) DownloadVideo(id string, path string) (bool, error) {

	fields := OptFields([]string{"embed.html"})

	options := []CallOption{fields}
	searchURL, err := addOptions(defaultBaseURL+"/"+id, options...)
	req, _ := http.NewRequest("GET", searchURL, nil)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", defaultOAuthHeader)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return false, err
	}

	//vm := otto.New()
	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var vidResp VideoResponse
	err = json.Unmarshal(data, &vidResp)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Parsing failed")
		return false, err
	}

	//doc, _ := html.Parse(strings.NewReader())
	htmlTokens := html.NewTokenizer(strings.NewReader(vidResp.Embed.HTML))

	var downloadURL string
	//var urlLink string
loop:
	for {
		tt := htmlTokens.Next()

		switch {
		case tt == html.ErrorToken:
			break loop
		case tt == html.StartTagToken:
			t := htmlTokens.Token()

			isAnchor := t.Data == "iframe"
			if isAnchor {
				for _, val := range t.Attr {
					if val.Key == "src" {
						downloadURL = val.Val
					}
				}
			}
		}
	}
	fmt.Println(downloadURL)
	req2, _ := http.NewRequest("GET", downloadURL, nil)
	q2 := req2.URL.Query()
	req2.URL.RawQuery = q2.Encode()
	req2.Header.Set("Authorization", defaultOAuthHeader)
	req2.Header.Set("Content-Type", "application/json")
	resp2, _ := c.client.Do(req2)

	data2, _ := ioutil.ReadAll(resp2.Body)
	defer resp2.Body.Close()
	//vm := otto.New()

	doc, err3 := goquery.NewDocumentFromReader(strings.NewReader(string(data2)))

	if err3 != nil {
		fmt.Println("Error parsing html 1")
		fmt.Println(err3)
	}

	if doc != nil {
		re := regexp.MustCompile("var config = ((?s).*?);")
		rm := re.FindStringSubmatch(doc.Find("script").Last().Text())
		if len(rm) != 0 {
			//fmt.Printf("%q", rm[1])
		}
		rmfmt := rm[1]
		var vimeoVideoConfig VimeoVideoConfig
		parseerr := json.Unmarshal([]byte(strings.TrimSpace(rmfmt)), &vimeoVideoConfig)
		if parseerr == nil {
			fmt.Println(vimeoVideoConfig.Request.Files.Progressive[0].URL)
		}
	}
	return true, nil
}

//GetVideoFromID retrieves the video details and returns the VideoDetails object
func GetVideoFromID(id string) (VideoDetails, error) {
	return VideoDetails{}, nil
}

//DownloadFileFromURL downloads the file defined by link argument to the specified path argument. Returns the number of bytes written
func DownloadFileFromURL(link string, path string) (int64, error) {
	out, patherr := os.Create(path)

	if patherr != nil {
		//TODO: Add debug log
		return 0, patherr
	}

	defer out.Close()

	resp, reqerr := http.Get(link)

	if reqerr != nil {
		//TODO: Add debug log
		return 0, reqerr
	}

	n, copyerr := io.Copy(out, resp.Body)
	defer resp.Body.Close()

	if copyerr != nil {
		return n, copyerr
	}
	return n, nil
}

//SearchAllVideos searches all the video matching the given querystr argument
func (c *Client) SearchAllVideos(querystr string) {

	query := OptQuery(strings.TrimSpace(querystr))
	uris := OptFields([]string{"uri"})
	//page := OptPage(2)
	//sort := OptSort("date")
	//direction := OptDirection("asc")
	for {
		//Add options here
		options := []CallOption{query, uris}
		searchURL, err := addOptions(defaultBaseURL, options...)

		req, _ := http.NewRequest("GET", searchURL, nil)
		q := req.URL.Query()
		req.URL.RawQuery = q.Encode()
		fmt.Println(req.URL.String())
		req.Header.Set("Authorization", defaultOAuthHeader)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.client.Do(req)

		if err != nil {
			fmt.Println(err.Error())
		}

		data, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			fmt.Println(err.Error)
		}

		fmt.Println(string(data))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

// CallOption is an optional argument to an API call.
// A CallOption is something that configures an API call in a way that is not specific to that API: page, filter and etc
type CallOption interface {
	Get() (key, value string)
}

// OptPage is an optional argument to an API call
type OptPage int

// Get return key/value for make query
func (o OptPage) Get() (string, string) {
	return "page", fmt.Sprint(o)
}

// OptPerPage is an optional argument to an API call
type OptPerPage int

// Get return key/value for make query
func (o OptPerPage) Get() (string, string) {
	return "per_page", fmt.Sprint(o)
}

// OptSort is an optional argument to an API call
type OptSort string

// Get key/value for make query
func (o OptSort) Get() (string, string) {
	return "sort", fmt.Sprint(o)
}

// OptDirection is an optional argument to an API call
// All sortable resources accept the direction parameter which must be either asc or desc.
type OptDirection string

// Get key/value for make query
func (o OptDirection) Get() (string, string) {
	return "direction", fmt.Sprint(o)
}

// OptFilter is an optional argument to an API call
type OptFilter string

// Get key/value for make query
func (o OptFilter) Get() (string, string) {
	return "filter", fmt.Sprint(o)
}

// OptFilterEmbeddable is an optional argument to an API call
type OptFilterEmbeddable string

// Get key/value for make query
func (o OptFilterEmbeddable) Get() (string, string) {
	return "filter_embeddable", fmt.Sprint(o)
}

// OptFilterPlayable is an optional argument to an API call
type OptFilterPlayable string

// Get key/value for make query
func (o OptFilterPlayable) Get() (string, string) {
	return "filter_playable", fmt.Sprint(o)
}

// OptQuery is an optional argument to an API call. Search query.
type OptQuery string

// Get key/value for make query
func (o OptQuery) Get() (string, string) {
	return "query", fmt.Sprint(o)
}

// OptFilterContentRating is an optional argument to an API call
// Content filter is a specific type of resource filter, available on all video resources.
// Any videos that do not match one of the provided ratings will be excluded from the list of videos.
// Valid ratings include: language/drugs/violence/nudity/safe/unrated
type OptFilterContentRating []string

// Get key/value for make query
func (o OptFilterContentRating) Get() (string, string) {
	return "filter_content_rating", strings.Join(o, ",")
}

// OptFields is an optional argument to an API call.
// With a simple parameter you can reduce the size of the responses,
// and dramatically increase the performance of your API requests.
type OptFields []string

// Get key/value for make query
func (o OptFields) Get() (string, string) {
	return "fields", strings.Join(o, ",")
}

// OptWeakSearch is an option argument to an API call
// to allow usage of legacy search on the vimeo backend to find private videos
type OptWeakSearch bool

// Get return key/value for make query
func (o OptWeakSearch) Get() (string, string) {
	return "weak_search", fmt.Sprint(o)
}

func addOptions(s string, opts ...CallOption) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs := u.Query()
	for _, o := range opts {
		qs.Set(o.Get())
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
