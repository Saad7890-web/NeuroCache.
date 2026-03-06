package kv

import "time"

func (s *Store) StartExpirationWorker() {
	ticker := time.NewTicker(1*time.Second)

	go func() {
		for range ticker.C {
			s.cleanupExpiredKeys()
		}
	}()
}

func (s *Store) cleanupExpiredKeys() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().Unix()

	for key, expireTime := range s.expiresAt {
		if now > expireTime {
			delete(s.data, key)
			delete(s.expiresAt, key)
		}
	}
}