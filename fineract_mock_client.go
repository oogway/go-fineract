package fineractor

import (
	"log"
	"net/http"
	"os"
	"path"
)

type MockTransport struct {
	DirectoryPath string
}

func (m *MockTransport) Do(req *http.Request) (*http.Response, error) {
	jsonResp, err := os.Open(path.Join(m.DirectoryPath, req.URL.Path+".json"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp := &http.Response{
		Body:       jsonResp,
		StatusCode: http.StatusOK,
	}
	return resp, nil

}
