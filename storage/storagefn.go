package storage

import "time"
import "strconv"
import "errors"

const sleeptimeInSec time.Duration = 4 * time.Second
const stdlifetime int = 3

//t must be key, or value
func IsExist(t string, v string) (bool, error) {
	if t == "key" {
		for _, s := range storage {
			if s.isKeyIn(v) {
				return true, nil
			}
		}
		return false, errors.New("The key doesn't exist")
	} else {
		for _, s := range storage {
		if s.isValueIn(v) {
			return true, nil
		}
	}
		return false, errors.New("Record doesn't have the value")
	}
}

func LifetimeManage() {
	for {
		time.Sleep(sleeptimeInSec)
		for _, s := range storage {
			if s.isZero() {
				DeleteStorageRecord(s.Key)
			} else {
				s.substructOne()
			}
		}
	}
}

func TestStorageAdding() {
	for i := 0; i < 3; i++ {
		AddStorageRecord(strconv.Itoa(i), "something")
	}
}

func AddStorageRecord(key, value string) error {
	exist, err := IsExist("key", key)
	err = errors.New("Record doesn't added (key isn't unique)")
	if exist {
		return err
	} else {
		storage[key] = &Storage{key, value, stdlifetime}
		return nil	
	}
}

func AddLifetime(key string) {
	storage[key].addLifetime()
}

func ReturnStorageRecord(key string) *Storage {
	return storage[key]
}

func ReturnKeyRecord(key string) string {
	return storage[key].Key
}

func ReturnValueRecord(key string) string {
	return storage[key].Value
}

func DeleteStorageRecord(key string) {
	delete(storage, key)
}

func ChangeStorageKey(oldk, newk string) error {
	err := AddStorageRecord(newk, ReturnValueRecord(oldk))
	if err != nil {
		return err
	}
	storage[newk].addLifetime()
	DeleteStorageRecord(oldk)
	return nil
}

func ChangeStorageValue(key, value string) {
	storage[key].setValue(value)
	storage[key].addLifetime()
}

func ReturnStorage() map[string]*Storage {
	return storage	
}