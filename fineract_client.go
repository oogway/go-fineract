package fineractor

import (
	"log"
	"net/http"
	"net/url"
	"sync"
)

type FineractOption struct{}

type Client struct {
	HostName   *url.URL
	Option     FineractOption
	HttpClient *http.Client
}

var once sync.Once
var client Client

func NewClient(hostName string, option FineractOption) (Fineractor, error) {
	host, err := url.Parse(hostName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	once.Do(func() {
		client = Client{
			HostName:   host,
			Option:     option,
			HttpClient: &http.Client{},
		}
	})
	return &client, err
}
