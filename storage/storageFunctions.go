package storage

import (
	"errors"
	"time"
)

const sleeptimeInSec time.Duration = 4 * time.Second

var storage = NewStorage()

func TestStorageAdding() {
	storage.Set("1", "something")
	storage.Set("2", "something new")
	storage.Set("3", "")
}

func NewStorage() *Storage {
	return &Storage {
		Record: make(map[string]*Record),
	}
}

func IsKeyExistReturnErr(key string) error {
	if storage.isKeyInStorage(key) {
		return nil
	}
	return errors.New("The key doesn't exist")
}

func IsValueNotNullReturnErr(key string) error {
	if storage.isValueInRecordNotNull(key) {
		return nil
	}
	return errors.New("Record doesn't have a value")
}

func IsKeyUniqueReturnErr(key string) error {
	if storage.isValueInRecordNotNull(key) {
		return errors.New("Record doesn't added (key isn't unique)")
	}
	return nil
}

func SetRecord(key, value string) {
	IsKeyExistReturnErr(key)
	storage.Set(key, value)
}

func AddRecordLifetime(key string) {
	IsKeyExistReturnErr(key)
	storage.AddRecordLifetime(key)
}

func ReturnStorageRecord(key string) *Record {
	IsKeyExistReturnErr(key)
	return storage.GetRecord(key)
}

func ReturnRecordValue(key string) string {
	return ReturnStorageRecord(key).Value
}

func DeleteStorageRecord(key string) {
	IsKeyExistReturnErr(key)
	delete(storage.Record, key)
}

func IsLifetimeZero(key string) bool {
	IsKeyExistReturnErr(key)
	return storage.IsLifetimeZero(key)
}

func SubstructLifetimeOne(key string) {
	IsKeyExistReturnErr(key)
	storage.SubstructLifetimeOne(key)
}

//high-order functions below

//func ReturnStorage() map[string]*Storage {
//	var
//}

func AddStorageRecord(key, value string) {
	IsKeyUniqueReturnErr(key)
	SetRecord(key, value)
	AddRecordLifetime(key)
}
//
func ChangeRecordValueReturnError(key, value string) {
	SetRecord(key, value)
	AddRecordLifetime(key)
}

func ChangeRecordKeyReturnErr(oldKey, newKey string) {
	AddStorageRecord(newKey, ReturnRecordValue(oldKey))
	AddRecordLifetime(newKey)
	DeleteStorageRecord(oldKey)
}

//goroutine
func LifetimeManage() {

	for {
		time.Sleep(sleeptimeInSec)
		for _, s := range storage.Record {
			SubstructLifetimeOne(s.getRecordKey())
			if IsLifetimeZero(s.getRecordKey()) {
				DeleteStorageRecord(s.getRecordKey())
			}
		}
	}
}
