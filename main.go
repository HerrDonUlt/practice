package main

import (
	strg "practicegit/storage"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	storage := strg.NewStorage()

	r := mux.NewRouter()

	//r.HandleFunc("/all", handlerShowAllRecords)
	r.HandleFunc("/{key}", HandlerShowRecord)

	r.HandleFunc("/setkey/{oldKey}/{newKey}", HandlerSetKey)
	r.HandleFunc("/changevalue/{key}/{value}", HandlerChangeValue)
	r.HandleFunc("/delete/{key}", HandlerDeleteRecord)
	r.HandleFunc("/value/{key}", HandlerShowValue)

	log.Fatal(http.ListenAndServe(":8080", r))

}
