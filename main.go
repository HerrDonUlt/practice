//split to packeges

package main

import (
	hdlr "practicegit/handlers"
	strt "practicegit/structs"
	"time"
)

const stdsleeptime time.Duration = 4
const stdlifetime int = 3

func testStorageAdding(stg map[string]*strt.KeyValInfo) {
	stg["1"] = &strt.KeyValInfo{"1", "something", stdlifetime}
	stg["2"] = &strt.KeyValInfo{"2", "something new", stdlifetime + 2}
	stg["3"] = &strt.KeyValInfo{"3", "", stdlifetime}
}

//substract the lifetime and delete zero lifetime KeyValInfo
func lifetimeManage(storage map[string]*strt.KeyValInfo) {
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

	var storage = make(map[string]*strt.KeyValInfo)

	hdlr.Storage = storage

	testStorageAdding(storage)

	go lifetimeManage(storage)

	hdlr.Handle()
}
