//incapsulation --im tried
//incapsulation
//golang race
//postman

package main

import (
	hdlr "practicegit/handlers"
	strg "practicegit/storage"
	// "fmt"
)

func main() {

	testStorageAdding()
	go strg.LifetimeManage()
	hdlr.HandleLoop()
}
