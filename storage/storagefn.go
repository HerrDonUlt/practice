package storage

import "time"
import "strconv"

const sleeptimeInSec time.Duration = 4 * time.Second
const stdlifetime int = 3

func IsKeyExist(k string) bool {
	for _, s := range storage {
		if s.IsKeyIn(k) {
			return true
		}
	}
	return false
}

func IsValueExist(k string) string {
	for _, s := range storage {
		if s.IsValueIn(k) {
			return ""
		}
	}
	return "Record doesn't have a value"
}

func LifetimeManage(storage map[string]*Storage) {
	for {
		time.Sleep(sleeptimeInSec)
		for _, s := range storage {
			if s.IsZero() {
				s.Delete(s.Key)
			} else {
				s.SubstructOne()
			}
		}
	}
}

func TestStorageAdding() {
	for i := 0; i < 3; i++ {
		storage[strconv.Itoa(i)].AddStorageRecord(strconv.Itoa(i), "something")
	}
}