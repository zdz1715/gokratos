package zmap

import "sync"

type StringMap[V any] struct {
	data map[string]V
	mu   sync.RWMutex
}

func (s *StringMap[V]) Set(key string, val V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.data == nil {
		s.data = make(map[string]V)
	}
	s.data[key] = val
}

func (s *StringMap[V]) Get(name string) (val V, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.data != nil {
		val, ok = s.data[name]
	}

	return
}

func (s *StringMap[V]) Removes(keys ...string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.data != nil {
		for _, key := range keys {
			delete(s.data, key)
		}
	}
}

func (s *StringMap[V]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func (s *StringMap[V]) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var (
		keys  = make([]string, len(s.data))
		index = 0
	)
	for key := range s.data {
		keys[index] = key
		index++
	}
	return keys
}

func (s *StringMap[V]) Values() []V {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var (
		values = make([]V, len(s.data))
		index  = 0
	)

	for _, value := range s.data {
		values[index] = value
		index++
	}
	return values
}
