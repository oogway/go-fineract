package fineract

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"sync"
)

const (
	fineractHost     = "13.209.34.65:8443" //"https://demo.openmf.org"
	fineractUser     = "mifos"
	fineractPassword = "password"
	baseURL          = "fineract-provider/api/v1/"
	Locale           = "en"
	DateFormat       = "dd MMMM yyyy"
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

func NewMockClient() (*Client, error) {
	return NewClient("https://"+fineractHost, fineractUser, fineractPassword, FineractOption{
		Transport: &MockTransport{DirectoryPath: "../testdata"},
	})
}

func clientsURL() string {
	return baseURL + "clients"
}

func paymentTypesURL() string {
	return baseURL + "paymenttypes"
}

func savingsAccountsURL() string {
	return baseURL + "savingsaccounts"
}

func headOfficeURL() string {
	return baseURL + "offices"
}
