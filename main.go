//incapsulation --im tried

package main
import (
	hdlr "practicegit/handlers"
	strt "practicegit/structs"
	// "fmt"
)
var storage = make(map[string]*strt.KeyValInfo)

//need auto testing
func testStorageAdding() {
	hdlr.AddStorageRecord("1", "something")
	hdlr.AddStorageRecord("2", "something new")
	hdlr.AddStorageRecord("3", "")
}

func main() {

	hdlr.Storage = storage

	testStorageAdding()
	go strt.LifetimeManage(storage)
	hdlr.HandleLoop()
}
