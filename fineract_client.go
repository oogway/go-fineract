package fineractor

import (
	"net/http"
	"sync"
)

type FineractOption struct{}

type Client struct {
	HostName   string
	Option     FineractOption
	HttpClient *http.Client
}

var once sync.Once
var client Client

func NewClient(hostName string, option FineractOption) Fineractor {
	httpClient := &http.Client{}
	once.Do(func() {
		client = Client{
			HostName:   hostName,
			Option:     option,
			HttpClient: httpClient,
		}
	})
	return &client
}
