package storage

import "time"

const sleeptimeInSec time.Duration = 4 * time.Second

func initTestStorage(s *Storage) {
	s.SetRecord("1", "something")
	s.SetRecord("2", "something new")
	s.SetRecord("3", "")
}

func NewStorage() *Storage {
	s := &Storage {records: make(map[string]*Record)}
	initTestStorage(s)
	go LifetimeManage(s)
	return s
}

func LifetimeManage(storage *Storage) {
	for {
		time.Sleep(sleeptimeInSec)
		for _, s := range storage.records {
			s.
			if IsLifetimeZero(s.getRecordKey()) {
				DeleteStorageRecord(s.getRecordKey())
			}
		}
	}
}
