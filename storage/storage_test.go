package storage

import "testing"

func TestInMemoryStorage(t *testing.T) {
	inMemoryStorage := NewInMemoryStorage()
	short := "000001"
	url := "http://example.com"
	inMemoryStorage.Save(short, url)
	retrievedURL, exists := inMemoryStorage.Get(short)
	if !exists {
		t.Error("expected retrieved URL to exist")
	}
	if retrievedURL != url {
		t.Errorf("expected retrieved URL to match the original URL, got: %s, expected: %s", retrievedURL, url)
	}
}

func TestGetNonExistentStorage(t *testing.T) {
	inMemoryStorage := InMemoryStorage{}
	retrievedURL, exists := inMemoryStorage.Get("nonexistent")
	if exists {
		t.Error("expected retrieved URL to not exist")
	}
	if retrievedURL != "" {
		t.Errorf("expected retrieved URL to be an empty string, got: %s", retrievedURL)
	}
}
