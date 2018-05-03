package fineractor

import "sync"

type FineractMockOption struct{}

type MockClient struct {
	DirectoryPath string
	Option        FineractMockOption
}

var mockOnce sync.Once
var mockClient MockClient

func NewMockClient(directoryPath string, option FineractMockOption) (Fineractor, error) {
	mockOnce.Do(func() {
		mockClient = MockClient{
			DirectoryPath: directoryPath,
			Option:        option,
		}
	})
	return &mockClient, nil
}
