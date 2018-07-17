//split to packeges

package main

import (
	h "practicegit/handlers"
	s "practicegit/structs"
	"time"
)

const stdsleeptime time.Duration = 4
const stdlifetime int = 3

func testStorageAdding(stg map[string]*s.KeyValInfo) {
	stg["1"] = &s.KeyValInfo{"1", "something", stdlifetime}
	stg["2"] = &s.KeyValInfo{"2", "something new", stdlifetime + 2}
	stg["3"] = &s.KeyValInfo{"3", "", stdlifetime}
}

//substract the lifetime and delete zero lifetime KeyValInfo
func lifetimeManage(storage map[string]*s.KeyValInfo) {
	for {
		time.Sleep(stdsleeptime * time.Second)
		for _, s := range storage {
			if s.LifeTime == 0 {
				delete(storage, s.Key)

			} else {
				s.LifeTime -= 1
			}
		}
	}
}

func main() {

	var storage = make(map[string]*s.KeyValInfo)

	h.Storage = storage

	testStorageAdding(storage)

	go lifetimeManage(storage)

	h.Handle()
}
