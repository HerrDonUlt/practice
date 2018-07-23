package main

import strg "practicegit/storage"
import hndl "practicegit/handlers"

func main() {
	go strg.LifetimeManage()
	hndl.InitHandlersAndStartServe()
}
