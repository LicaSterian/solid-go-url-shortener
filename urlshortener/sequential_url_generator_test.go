package urlshortener

import (
	"solid-go-url-shortener/storage"
	"testing"
)

func TestSequentialGenerator(t *testing.T) {
	mockStorage := storage.NewMockStorage()
	sequentialGenerator := &SequentialURLGenerator{}
	service := NewURLShortener(mockStorage, sequentialGenerator)

	short1 := service.Shorten("https://example.com/1")
	short2 := service.Shorten("https://example.com/2")

	if short1 == short2 {
		t.Errorf("expected different short URLs, got '%s' and '%s'", short1, short2)
	}
}
