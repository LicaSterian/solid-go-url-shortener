package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

type URLShortener struct {
	mapping map[string]string
	mu      sync.Mutex
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		mapping: make(map[string]string),
	}
}

func (u *URLShortener) Shorten(url string) string {
	u.mu.Lock()
	defer u.mu.Unlock()

	short := fmt.Sprintf("%06d", rand.Intn(1000000))
	u.mapping[short] = url
	return short
}

func (u *URLShortener) Get(short string) (string, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	url, exists := u.mapping[short]
	if !exists {
		return "", errors.New("short URL not found")
	}
	return url, nil
}

func main() {
	service := NewURLShortener()

	short := service.Shorten("https://example.com")
	fmt.Println("Shortened URL:", short)

	original, err := service.Get(short)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Original URL:", original)
	}
}
