package main

import strg "practicegit/storage"
import hndl "practicegit/handlers"

func main() {
	var storage = strg.NewStorage()
	go strg.LifetimeManage(storage)
	hndl.InitHandlers(storage)
}
