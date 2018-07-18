//incapsulation --im tried

package main

import (
	hdlr "practicegit/handlers"
	strt "practicegit/structs"
	"time"
)

const stdlifetime int = 3

//need auto testing
func testStorageAdding(stg map[string]*strt.KeyValInfo) {
	stg["1"] = &strt.KeyValInfo{"1", "something", stdlifetime}
	stg["2"] = &strt.KeyValInfo{"2", "something new", stdlifetime + 2}
	stg["3"] = &strt.KeyValInfo{"3", "", stdlifetime}
}

func main() {

	var storage = make(map[string]*strt.KeyValInfo)

	hdlr.Storage = storage

	testStorageAdding(storage)
	go strt.LifetimeManage(storage)

	hdlr.Handle()
}
