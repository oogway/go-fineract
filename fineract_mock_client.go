package fineractor

type FineractMockOption struct{}

type MockClient struct {
	FolderPath string
	Option     FineractMockOption
}

func NewMockClient(folderPath string, option FineractMockOption) *MockClient {
	return &MockClient{
		FolderPath: folderPath,
		Option:     option,
	}
}
