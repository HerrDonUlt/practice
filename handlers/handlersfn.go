package handlers

import strt "practicegit/structs"

var Storage map[string]*strt.KeyValInfo

func extendLifetimeFn(s map[string]*strt.KeyValInfo, key string) {
	s[key].AddLifetime(key)
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