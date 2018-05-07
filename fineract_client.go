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
	UserName   string
	Password   string
	Option     FineractOption
	HttpClient *http.Client
}

var once sync.Once
var client Client

func NewClient(hostName, userName, password string, option FineractOption) (Fineractor, error) {
	host, err := url.Parse(hostName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	once.Do(func() {
		client = Client{
			HostName:   host,
			UserName:   userName,
			Password:   password,
			Option:     option,
			HttpClient: &http.Client{},
		}
	})
	return &client, err
}
