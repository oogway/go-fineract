package fineractor

type FineractMockOption struct{}

type MockClient struct {
	DirectoryPath string
	Option        FineractMockOption
}

func NewMockClient(directoryPath string, option FineractMockOption) Fineractor {
	return MockClient{
		DirectoryPath: directoryPath,
		Option:        option,
	}
}
