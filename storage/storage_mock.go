package storage

import "fmt"

type MockStorage struct {
	mapping map[string]string
}

func NewMockStorage() *MockStorage {
	fmt.Println("NewMockStorage")
	return &MockStorage{mapping: make(map[string]string)}
}

func (m *MockStorage) Save(short, url string) {
	m.mapping[short] = url
}
func (m *MockStorage) Get(short string) (string, bool) {
	url, exists := m.mapping[short]
	return url, exists
}
