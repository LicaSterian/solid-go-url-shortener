package urlshortener

import (
	"solid-go-url-shortener/storage"
	"testing"
)

func TestShorten(t *testing.T) {
	mockStorage := storage.NewMockStorage()
	mockGenerator := &MockGenerator{value: "123456"}
	service := NewURLShortener(mockStorage, mockGenerator)

	short := service.Shorten("https://example.com")
	if short != "123456" {
		t.Errorf("expected short URL '123456', got '%s'", short)
	}

	url, err := service.Get(short)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://example.com" {
		t.Errorf("expected original URL 'https://example.com', got '%s'", url)
	}
}

func TestGetNonExistent(t *testing.T) {
	mockStorage := storage.NewMockStorage()
	mockGenerator := &MockGenerator{}
	service := NewURLShortener(mockStorage, mockGenerator)

	_, err := service.Get("nonexistent")
	if err == nil {
		t.Fatal("expected error for non-existent short URL, got nil")
	}
}
