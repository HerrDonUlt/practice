package main

import (
	"fmt"
)

type Storage struct {
	Pair map[string]string
	LifeTime int    `json:"life_time"`
}

func TestStorageAdding(storage *Storage) {
	storage.set("1", "something")
	storage.set("2", "something new")
	storage.set("3", "1")
}


func NewStorage() *Storage {
	return &Storage {
		Pair: make(map[string]string),
	}
}

func (s *Storage) set(key, value string) {
	s.Pair[key] = value
}

func main()  {
	srg := NewStorage()

	TestStorageAdding(srg)

	fmt.Println(srg)

	func(key string) {
		for k := range srg.Pair {
			if k == key {
				fmt.Println("t")
			}
		}
	}("4")
}

