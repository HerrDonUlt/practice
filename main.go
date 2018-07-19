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

	strg.TestStorageAdding()
	go strg.LifetimeManage()
	hdlr.HandleLoop()
}
