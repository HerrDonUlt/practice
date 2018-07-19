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

func LifetimeManage(storage map[string]*Storage) {
	for {
		time.Sleep(sleeptimeInSec)
		for _, s := range storage {
			if s.isZero() {
				s.delete(s.Key)
			} else {
				s.substructOne()
			}
		}
	}
}

func TestStorageAdding() {
	for i := 0; i < 3; i++ {
		storage[strconv.Itoa(i)].addStorageRecord(strconv.Itoa(i), "something")
	}
}

// func AddStorageRecord(key string, value string) {
// 	exist, err := IsExist("key", key)
// 	{
// 		storage[key].addStorageRecord(key, value)
// 	}
// }