package urlshortener

import (
	"fmt"
	"solid-go-url-shortener/storage"
)

type URLShortenerGenerator interface {
	Shorten(url string) string
}

type URLShortenerRetriever interface {
	Get(short string) (string, error)
}

type URLShortener interface {
	URLShortenerGenerator
	URLShortenerRetriever
}

type Shortener struct {
	storage   storage.Storage
	generator URLGenerator
}

func NewURLShortener(storage storage.Storage, generator URLGenerator) URLShortener {
	return &Shortener{storage: storage, generator: generator}
}

func (u *Shortener) Shorten(url string) string {
	short := u.generator.Generate()
	u.storage.Save(short, url)
	return short
}

func (u *Shortener) Get(short string) (string, error) {
	url, exists := u.storage.Get(short)
	if !exists {
		return "", fmt.Errorf("short URL not found")
	}
	return url, nil
}
