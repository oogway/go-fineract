package fineract

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"sync"
)

const (
	fineractHost     = "https://13.209.34.65:8443" //"https://demo.openmf.org"
	fineractUser     = "mifos"
	fineractPassword = "password"
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

func NewMockClient(mockTransport *MockTransport) (*Client, error) {
	if mockTransport == nil {
		mockTransport = &MockTransport{DirectoryPath: "../testdata"}
	}
	return NewClient(fineractHost, fineractUser, fineractPassword, FineractOption{
		Transport: mockTransport,
	})
}
