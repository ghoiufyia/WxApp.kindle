package session

import (
	"sync"
	"encoding/json"
	// "fmt"
)

type Store struct {
	sid		string
	lock	sync.RWMutex
	values	map[string]interface{}
	handle	handler
}

func (s *Store)GetSid() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.sid
}

func (s *Store)Save() error {
	s.lock.Lock()
	defer s.lock.Unlock()

	jsonStr,err := json.Marshal(s.values)
	if err != nil {
		return err
	}
	s.handle.Write(s.sid,string(jsonStr))
	return nil
}

func (s *Store)Set(key string,value interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.values[key] = value
	return nil
}

func (s *Store)Get(key string) interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if v, ok := s.values[key]; ok {
		return v
	}
	return nil
}

func (s *Store)Delete(key string) error  {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.values, key)
	return nil
}

func (s *Store)Flush() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.values = make(map[string]interface{})


	return nil
}

