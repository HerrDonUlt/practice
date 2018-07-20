//golang race
//postman

//переделать ретернсторадж
//peredelat storage

package main

import (
	hdlr "practicegit/handlers"
	strg "practicegit/storage"

)

func main() {
	strg.TestStorageAdding()
	go strg.LifetimeManage()
	hdlr.HandleEstablish()
}
