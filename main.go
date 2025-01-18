package main

import (
	"fmt"

	"solid-go-url-shortener/storage"
	"solid-go-url-shortener/urlshortener"
)

type URLShortener struct {
	storage   storage.Storage
	generator urlshortener.URLGenerator
}

func NewURLShortener(storage storage.Storage, generator urlshortener.URLGenerator) *URLShortener {
	return &URLShortener{storage: storage, generator: generator}
}

func (u *URLShortener) Shorten(url string) string {
	short := u.generator.Generate()
	u.storage.Save(short, url)
	return short
}

func (u *URLShortener) Get(short string) (string, error) {
	url, exists := u.storage.Get(short)
	if !exists {
		return "", fmt.Errorf("short URL not found")
	}
	return url, nil
}

func main() {
	inMemoryStorage := storage.NewInMemoryStorage()
	generator := &urlshortener.RandomURLGenerator{}
	service := NewURLShortener(inMemoryStorage, generator)

	short := service.Shorten("https://example.com")
	fmt.Println("Shortened URL:", short)

	original, err := service.Get(short)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Original URL:", original)
	}
}
