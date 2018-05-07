package fineractor

type FineractMockOption struct{}

type MockClient struct {
	DirectoryPath string
	Option        FineractMockOption
}

func NewMockClient(directoryPath string, option FineractMockOption) (Fineractor, error) {
	mockClient := MockClient{
		DirectoryPath: directoryPath,
		Option:        option,
	}
	return &mockClient, nil
}
