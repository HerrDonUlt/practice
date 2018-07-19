package storage

import "sync"

var storage = make(map[string]*Storage)

//store the info about key-val thing
type Storage struct {
	sync.Mutex
	Key      string `json:"key"`
	Value    string `json:"value"`
	LifeTime int    `json:"life_time"`
}

func (s *Storage) substructLifetimeOne() {
	s.Lock()
	defer s.Unlock()
	s.LifeTime -= 1
}

func (s Storage) isLifetimeZero() bool {
	s.Lock()
	defer s.Unlock()
	if s.LifeTime == 0 {
		return true
	} 
	return false
}

func (s Storage) isKeyIn(k string) bool {
	s.Lock()
	defer s.Unlock()
	if s.Key == k {
		return true
	}
	return false
}

func (s Storage) isValueIn(k string) bool {
	s.Lock()
	defer s.Unlock()
	if s.Key == k && s.Value != "" {
		return true
	}
	return false
}

func (s *Storage) setKey(k string) {
	s.Lock()
	defer s.Unlock()
	s.Key = k
}

func (s *Storage) setValue(v string) {
	s.Lock()
	defer s.Unlock()
	s.Value = v
}

func (s *Storage) addRecordLifetime() {
	s.Lock()
	defer s.Unlock()
	s.LifeTime += stdlifetime
}
