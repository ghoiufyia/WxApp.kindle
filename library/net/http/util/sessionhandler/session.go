package sessionhandler

import (
	// "sync"
	// "encoding/json"
	// "fmt"
)

// type Session struct {
// 	Sid		string
// 	lock	sync.RWMutex
// 	Values	map[string]interface{}
// }


// func (s *Session)GetSid() string {
// 	s.lock.RLock()
// 	defer s.lock.RUnlock()
// 	return s.sid
// }

// func (s *Session)Save() error {
// 	s.lock.Lock()
// 	defer s.lock.Unlock()

// 	jsonStr,err := json.Marshal(s.values)
// 	if err != nil {
// 		return err
// 	}
// 	s.handle.Write(s.sid,string(jsonStr))
// 	return nil
// }

// func (s *Session)Set(key string,value interface{}) error {
// 	s.lock.Lock()
// 	defer s.lock.Unlock()
// 	s.values[key] = value
// 	return nil
// }

// func (s *Session)Get(key string) interface{} {
// 	s.lock.RLock()
// 	defer s.lock.RUnlock()
// 	if v, ok := s.values[key]; ok {
// 		return v
// 	}
// 	return nil
// }

// func (s *Session)Delete(key string) error  {
// 	s.lock.Lock()
// 	defer s.lock.Unlock()
// 	delete(s.values, key)
// 	return nil
// }

// func (s *Session)Flush() error {
// 	s.lock.Lock()
// 	defer s.lock.Unlock()
// 	s.values = make(map[string]interface{})


// 	return nil
// }

