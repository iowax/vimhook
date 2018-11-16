package vimeo

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//Client struct to handle communication with vimeo api
type Client struct {
	client *http.Client
	dbconn *DBConnector
}

//NewVimeoClient provides an instance of a VimeoClient to connect to vimeo API
func NewVimeoClient() (*Client, error) {
	restclient := http.DefaultClient
	dbinstance, err := NewDBConnector(defaultMgoURL)

	if err != nil {
		return &Client{client: restclient, dbconn: dbinstance}, err
	}
	return nil, err
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
