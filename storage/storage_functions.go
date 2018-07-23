package storage

import (
	"time"
)

const sleeptimeInSec time.Duration = 4 * time.Second

func initTestStorage(s *Storage) {
	s.SetRecord("1", "something")
	s.SetRecord("2", "something new")
	s.SetRecord("3", "")
}

func NewStorage() *Storage {
	s := &Storage{records: make(map[string]*Record)}
	initTestStorage(s)

	return s
}

func LifetimeManage(storage *Storage) {
	for {
		
		time.Sleep(sleeptimeInSec)
		storage.DeleteNullStorageRecords()
		storage.SubstructLifetimeRecords()
	}
}
