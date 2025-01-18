package storage

import "testing"

// I wrote this test because the coverage does not include the usage of MockStorage from other packages
func TestMockStorage(t *testing.T) {
	mockStorage := NewMockStorage()
	short := "000001"
	url := "http://example.com"
	mockStorage.Save(short, url)
	retrievedURL, exists := mockStorage.Get(short)
	if !exists {
		t.Error("expected mock storage url to exist")
	}
	if retrievedURL != url {
		t.Errorf("expected retrieved URL to match actual URL, got: %s, want: %s", retrievedURL, url)
	}
}
