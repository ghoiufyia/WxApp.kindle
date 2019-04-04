package session

import (
	"sync"
)

type Store struct {
	sid		string
	lock	sync.RWMutex
	values	map[interface{}]interface{}
}

func (s *Store)Set(key,value interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.values[key] = value
	return nil
}

func (s *Store)Get(key interface{}) interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if v, ok := s.values[key]; ok {
		return v
	}
	return nil
}

func (s *Store)Delete(key interface{}) error  {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.values, key)
	return nil
}

func (s *Store)Flush() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.values = make(map[interface{}]interface{})
	return nil
}

func (s *Store)Save() error {
	
	return nil
}

