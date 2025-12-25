package storage

import (
	"sync"
	"time"
)

type Store struct {
	mu   sync.RWMutex
	data map[string]Item
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]Item),
	}
}

func (s *Store) Set(key, value string, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item := Item{
		Value: value,
	}

	if ttl > 0 {
		item.HasExpiry = true
		item.ExpireAt = time.Now().Add(ttl)
	}

	s.data[key] = item
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	item, exists := s.data[key]
	s.mu.RUnlock()

	if !exists {
		return "", false
	}

	if item.HasExpiry && time.Now().After(item.ExpireAt) {
		s.mu.Lock()
		delete(s.data, key)
		s.mu.Unlock()
		return "", false
	}

	return item.Value, true
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	delete(s.data, key)
	s.mu.Unlock()
}
