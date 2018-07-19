package storage

import "time"
import (
	"errors"
	"sync"
)

const sleeptimeInSec time.Duration = 4 * time.Second
const stdlifetime int = 3

func TestStorageAdding() {
	s.Lock()
	defer s.Unlock()
	storage["1"] = &Storage{sync.Mutex{}, "1", "something", stdlifetime}
}

//t must be key, or value
func IsKeyExist(key string) error {
	s.Lock()
	defer s.Unlock()
	for _, s := range storage {
		if s.isKeyIn(key) {
			return nil
		}
	}
	return errors.New("The key doesn't exist")
}

func IsValueExist(key string) error {
	s.Lock()
	defer s.Unlock()
	for _, s := range storage {
		if s.isValueIn(key) {
			return nil
		}
	}
	return errors.New("Record doesn't have the value")
}

func LifetimeManage() {
	for {
		time.Sleep(sleeptimeInSec)
		for _, s := range storage {
			if s.isLifetimeZero() {
				DeleteStorageRecord(s.Key)
			} else {
				s.substructLifetimeOne()
			}
		}
	}
}

func AddStorageRecord(key, value string) error {
	err := IsKeyExist(key)
	if err != nil {
		return errors.New("Record doesn't added (key isn't unique)")
	}
	storage[key].setKey(key)
	storage[key].setValue(value)
	storage[key].addRecordLifetime()
	return nil
}

func AddRecordLifetime(key string) {
	storage[key].addRecordLifetime()
}

func ReturnStorageRecord(key string) *Storage {
	return storage[key]
}

func ReturnRecordKey(key string) string {
	return storage[key].Key
}

func ReturnRecordValue(key string) string {
	return storage[key].Value
}

func DeleteStorageRecord(key string) {
	delete(storage, key)
}

func ChangeRecordKey(oldKey, newKey string) error {
	err := AddStorageRecord(newKey, ReturnRecordValue(oldKey))
	if err != nil {
		return err
	}
	storage[newKey].addRecordLifetime()
	DeleteStorageRecord(oldKey)
	return nil
}

func ChangeRecordValue(key, value string) {
	storage[key].setValue(value)
	storage[key].addRecordLifetime()
}

//func ReturnStorage() map[string]*Storage {
//	var
//}
