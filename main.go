

package main

import (
	h "practicegit/handlers"
	s "practicegit/structs"
)



const stdlifetime int = 3

func testStorageAdding(stg map[string]*s.KeyValInfo) {
	stg["1"] = &s.KeyValInfo{"1", "something", stdlifetime}
	stg["2"] = &s.KeyValInfo{"2", "something new", stdlifetime + 2}
	stg["3"] = &s.KeyValInfo{"3", "", 0}
}

func main() {

	var storage = make(map[string]*s.KeyValInfo)

	h.Storage = storage

	testStorageAdding(storage)

	go s.LifetimeManage(storage)

	h.Handle()
}
