package fineract

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

type MockTransport struct {
	DirectoryPath string
}

func (m *MockTransport) Do(req *http.Request) (*http.Response, error) {
	if req.Method != http.MethodGet {
		filePath := path.Join(m.DirectoryPath, fmt.Sprintf("%s_%s.json", req.URL.Path, req.Method))
		response, err := m.getResponseFromFile(filePath)
		if err == nil {
			return response, nil
		}
	}
	filePath := path.Join(m.DirectoryPath, req.URL.Path+".json")
	return m.getResponseFromFile(filePath)
}

func (m *MockTransport) getResponseFromFile(filePath string) (*http.Response, error) {
	jsonResp, err := os.Open(filePath)
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

func random(min, max int) uint64 {
	rand.Seed(time.Now().UnixNano())
	return uint64(rand.Intn(max-min) + min)
}
