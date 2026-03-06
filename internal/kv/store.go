package kv

import (
	"sync"
	"time"
)

type Store struct {
	data map[string]string
	expiresAt map[string]int64
	mu sync.RWMutex
}


func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
		expiresAt: make(map[string]int64),
	}
}

func (s *Store) Set(key, value string, ttlSeconds int){
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	if ttlSeconds > 0 {
		s.expiresAt[key] = time.Now().Unix() + int64(ttlSeconds)
	}
}


func (s *Store) Get(key string) (string, bool){
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.isExpired(key) {
		delete(s.data, key)
		delete(s.expiresAt, key)
		return "", false
	}


	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Del(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_,ok := s.data[key]
	if ok {
		delete(s.data, key)
	}
	return ok
}

func (s *Store) isExpired(key string) bool {

	expireTime, ok := s.expiresAt[key]

	if !ok {
		return false
	}

	return time.Now().Unix() > expireTime
}