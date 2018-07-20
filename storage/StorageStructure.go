package storage

import "sync"

type Storage struct {
	sync.Mutex
	Record map[string]*Record
}

//-- move
func (s *Storage) count() int {
	return len(s.Record)
}

func (s *Storage) isKeyInStorage(key string) bool {
	for k := range s.Record {
		if k == key {
			return true
		}
	}
	return false
}

//-- high-order methods below

func (s *Storage) getRecord(key string) *Record {
	if s.count() > 0 {
		return s.Record[key]
	}
	return nil
}

func (s *Storage) getRecordValue(key string) string {
	return s.Record[key].Value
}

func (s *Storage) isValueInRecordNotNull(key string) bool {
	if s.isKeyInStorage(key) && s.getRecordValue(key) != "" {
		return true
	}
	return false
}

func (s *Storage) substructLifeTimeOne(key string) {
	s.Record[key].LifeTime -= 1
}
