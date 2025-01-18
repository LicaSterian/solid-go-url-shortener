package main

import (
	"fmt"

	"solid-go-url-shortener/storage"
	"solid-go-url-shortener/urlshortener"
)

func main() {
	inMemoryStorage := storage.NewInMemoryStorage()
	sequentialURLGenerator := &urlshortener.SequentialURLGenerator{}
	service := urlshortener.NewURLShortener(inMemoryStorage, sequentialURLGenerator)

	short := service.Shorten("https://example.com")
	fmt.Println("Shortened URL:", short)

	original, err := service.Get(short)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Original URL:", original)
	}
}
