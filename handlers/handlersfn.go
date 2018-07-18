package handlers

import strt "practicegit/structs"

const stdlifetime int = 3

var Storage map[string]*strt.KeyValInfo

func AddStorageRecord(key string, value string) {
	Storage[key] = &strt.KeyValInfo{key,value, stdlifetime}
}

func DeleteStorageRecord(key string) {
	delete(Storage, key)
}

func AddLifetime(key string) {
	Storage[key].LifeTime += stdlifetime
}

func isKeyExist(k string) string {
	for _, s := range Storage {
		if s.IsKeyIn(k) {
			return ""
		}
	}
	return "this key doesn't exist"
}

func isValueExist(k string) string {
	for _, s := range Storage {
		if s.IsValueIn(k) {
			return ""
		}
	}
	return "Record doesn't have a value"
}