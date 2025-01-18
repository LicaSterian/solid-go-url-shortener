package storage

import "sync"

type Storage interface {
	Save(short string, url string)
	Get(short string) (string, bool)
}

type InMemoryStorage struct {
	mapping map[string]string
	mu      sync.Mutex
}

func NewInMemoryStorage() Storage {
	return &InMemoryStorage{
		mapping: make(map[string]string),
	}
}

func (s *InMemoryStorage) Save(short, url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.mapping[short] = url
}

func (s *InMemoryStorage) Get(short string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	url, exists := s.mapping[short]
	return url, exists
}
