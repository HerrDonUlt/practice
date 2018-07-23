package storage

import (
	"sync"
)

const stdlifetime int = 3

var Strg Storage

func init() {
	Strg = NewStorage()
}

type Record struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	LifeTime int    `json:"life_time"`
}

type Storage struct {
	mux     sync.RWMutex
	records map[string]*Record
}

func (s Storage) IsKeyInStorage(key string) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	for k := range s.records {
		if k == key {
			return true
		}
	}
	return false
}

func (r Record) IsValueInRecordNotNull(key string) bool {
	if r.Value != "" {
		return true
	}
	return false
}

func (s Storage) GetRecord(key string) *Record {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.records[key]
}

func (s Storage) GetAllRecord() map[string]*Record {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.records
}

func (s Storage) SetRecord(key, value string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.records[key] = &Record{Key: key, Value: value, LifeTime: stdlifetime}
}

func (r Record) getRecordValue() string {
	return r.Value
}

func (r Record) getRecordKey() string {
	return r.Key
}

func (s Storage) SubstructLifetimeRecords() {
	s.mux.Lock()
	defer s.mux.Unlock()
	for _, r := range s.records {
		r.LifeTime -= 1
	}
}

func (s Storage) DeleteNullStorageRecords() {
	s.mux.Lock()
	defer s.mux.Unlock()
	for _, r := range s.records {
		if r.LifeTime == 0 {
			delete(s.records, r.Key)
		}
	}
	return
}
