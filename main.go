//incapsulation --im tried

package main

import (
	hdlr "practicegit/handlers"
	strt "practicegit/structs"
	"fmt"
)

const stdlifetime int = 3

//need auto testing
func testStorageAdding(stg map[string]*strt.KeyValInfo) {
	stg["1"] = &strt.KeyValInfo{"1", "something", 10 }
	// stg["1"].AddRecord("1", "something")
	// stg["2"].AddRecord("2", "something new")
	// stg["3"].AddRecord("3", "")
}

func main() {

	var storage = make(map[string]*strt.KeyValInfo)

	hdlr.Storage = storage

	testStorageAdding(storage)
	go strt.LifetimeManage(storage)
	// fmt.Println(storage["1"])

	hdlr.Handle()
}
