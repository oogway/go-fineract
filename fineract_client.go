package fineract

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"sync"
)

type Transporter interface {
	Do(req *http.Request) (*http.Response, error)
}

type FineractOption struct {
	Transport  Transporter
	SkipVerify bool
}

type Client struct {
	HostName *url.URL
	UserName string
	Password string
	Option   FineractOption
}

var once sync.Once
var client Client

func NewClient(hostName, userName, password string, option FineractOption) (*Client, error) {
	host, err := url.Parse(hostName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	once.Do(func() {
		if option.Transport == nil {
			httpClient := http.Client{}
			if option.SkipVerify == true {
				httpClient.Transport = &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				}
			}
			option.Transport = &httpClient
		}

		client = Client{
			HostName: host,
			UserName: userName,
			Password: password,
			Option:   option,
		}
	})
	return &client, err
}

func NewMockClient() (*Client, error){
	return NewClient("https://demo.openmf.org", "mifos", "password", FineractOption{
		Transport: &MockTransport{DirectoryPath: "../testdata"},
	})
}
