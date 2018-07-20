package main

import (
	hdlr "practicegit/handlers"
	strg "practicegit/storage"

	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main() {
	storage := strg.NewStorage()

	r := mux.NewRouter()

	//r.HandleFunc("/all", handlerShowAllRecords)
	r.HandleFunc("/{key}", handlerShowRecord)

	r.HandleFunc("/setkey/{oldKey}/{newKey}", handlerSetKey)
	r.HandleFunc("/changevalue/{key}/{value}", handlerChangeValue)
	r.HandleFunc("/delete/{key}", handlerDeleteRecord)
	r.HandleFunc("/value/{key}", handlerShowValue)

	log.Fatal(http.ListenAndServe(":8080", r))

}
