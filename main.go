package main

import (
	strg "practicegit/storage"
	hndl "practicegit/handlers"
)

func main() {
	strg.InitTestStorage()
	go strg.LifetimeManage()
	hndl.InitHandlersAndStartServe()
}
