package storage

import "time"
import (
	"errors"
)

const sleeptimeInSec time.Duration = 4 * time.Second
const stdlifetime int = 3

func TestStorageAdding() {
	storage["1"].Record.setKey("1")
	storage["1"].Record.setValue("something")
	storage["1"].addRecordLifetime()
}

//t must be key, or value
func IsKeyExist(key string) error {
	for _, s := range storage {
		if s.Record.isKeyIn(key) {
			return nil
		}
	}
	return errors.New("The key doesn't exist")
}

func IsValueExist(key string) error {
	for _, s := range storage {
		if s.Record.isValueIn(key) {
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
				DeleteStorageRecord(s.Record.Key)
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
	storage[key].Record.setKey(key)
	storage[key].Record.setValue(value)
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
	return storage[key].Record.Key
}

func ReturnRecordValue(key string) string {
	return storage[key].Record.Value
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
	storage[key].Record.setValue(value)
	storage[key].addRecordLifetime()
}

//func ReturnStorage() map[string]*Storage {
//	var
//}
