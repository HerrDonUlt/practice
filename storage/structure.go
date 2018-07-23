package storage

import (
	"sync"
)

const stdlifetime int = 3

type Record struct {
	Key      string `json:"key,omitempty"`
	Value    string `json:"value,omitempty"`
	LifeTime int    `json:"life_time"`
}

var storage = struct {
	sync.RWMutex
	records map[string]Record
}{ records: make(map[string]Record), }

func IsKeyInStorage(key string) bool {
	storage.RLock()
	defer storage.RUnlock()
	for k := range storage.records {
		if k == key {
			return true
		}
	}
	return false
}

func IsValueInStorageNotNull(key string) bool {
	storage.RLock()
	defer storage.RUnlock()
	if storage.records[key].Value != "" {
		return true
	}
	return false
}

func GetRecord(key string) Record {
	storage.Lock()
	defer storage.Unlock()
	return storage.records[key]
}

func GetRecordValue(record Record) string {
	return record.Value
}

func GetAllRecord() map[string]Record {
	storage.RLock()
	defer storage.RUnlock()
	return storage.records
}

func SetRecord(key, value string) {
	storage.Lock()
	defer storage.Unlock()
	storage.records[key] = Record{Key: key, Value: value, LifeTime: stdlifetime}
}

func SubstructLifetimeRecords() {
	storage.Lock()
	defer storage.Unlock()
	for key, _ := range storage.records {
		r := storage.records[key]
		r.LifeTime -= 1
		storage.records[key] = r
	}
}

func DeleteNullStorageRecords() {
	storage.Lock()
	defer storage.Unlock()
	for _, r := range storage.records {
		if r.LifeTime == 0 {
			delete(storage.records, r.Key)
		}
	}
}

func AddLifetimeForRecord(key string) {
	storage.Lock()
	defer storage.Unlock()
	r := storage.records[key]
	r.LifeTime += stdlifetime
	storage.records[key] = r
}

func DeleteStorageRecord(key string) {
	storage.Lock()
	defer storage.Unlock()
	delete(storage.records, key)
}
