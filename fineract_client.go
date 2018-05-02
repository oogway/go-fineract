package fineractor

type FineractOption struct{}

type Client struct {
	HostName string
	Option   FineractOption
}

func NewClient(hostName string, option FineractOption) *Client {
	return &Client{
		HostName: hostName,
		Option:   option,
	}
}
