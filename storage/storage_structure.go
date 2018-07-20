package storage

import "sync"

type Record struct {
	Key string `json:"key"`
	Value string `json:"value"`
	LifeTime int    `json:"life_time"`
}

type Storage struct {
	mux sync.RWMutex
	records map[string]*Record
}

func (s *Storage) IsKeyInStorage(key string) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	for k := range s.records {
		if k == key {
			return true
		}
	}
	return false
}

func (r *Record) IsValueInRecordNotNull(key string) bool {
	if r.Value != "" {
		return true
	}
	return false
}

func (s *Storage) GetRecord(key string) *Record {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.records[key]
}

func (s *Storage) SetRecord(key, value string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.records[key].Key = key
	s.records[key].Value = value
}

func (r *Record) GetRecordValue(key string) string {
	return r.Value
}

func (s *Storage) SubstructLifeTimeOne(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.records[key].LifeTime -= 1
}
