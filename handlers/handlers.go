package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"practicegit/storage"
)

func handlerShowRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	record := storage.Strg.GetRecord(vars["key"])
	json.NewEncoder(w).Encode(record)
	log.Println("Record " + vars["key"] + " showed")
}

func handlerShowAllRecords(w http.ResponseWriter, r *http.Request) {
	records := storage.Strg.GetAllRecord()
	for _, s := range records {

		json.NewEncoder(w).Encode(s)

	}
	log.Println("showed all records")
}

//
//func HandlerShowValue(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	record := strg.GetRecord(vars["key"])
//	json.NewEncoder(w).Encode(record)
//	log.Println("Value " + vars["key"] + " showed")
//}

//re
//func HandlerSetKey(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	if strg.IsKeyInStorage(vars["key"]) {
//
//	}
//	if err != nil {
//		strg.AddStorageRecord(vars["oldKey"], "")
//		encoder.setMessage("New record seted")
//		encoder.encodeMessage(w)
//		log.Println(encoder.ActMessage)
//		return
//	}
//	err = strg.IsKeyExistReturnErr("newKey")
//	if err != nil {
//		encoder.setError("New key already in use")
//		encoder.encodeValue(w, vars["key"])
//		log.Println(encoder.ErrMessage)
//		return
//	}
//	strg.ChangeRecordKeyReturnErr(vars["oldKey"], vars["newKey"])
//	encoder.setMessage("Key changed")
//	encoder.encodeMessage(w)
//	log.Println(encoder.ActMessage)
//}
//
//func HandlerChangeValue(w http.ResponseWriter, r *http.Request) {
//	var encoder Encoder
//	vars := mux.Vars(r)
//
//	err := strg.IsKeyExistReturnErr(vars["key"])
//	if err != nil {
//		encoder.setError("Key is not exist")
//		encoder.encodeError(w)
//		log.Println(encoder.ErrMessage)
//		return
//	}
//	strg.ChangeRecordValueReturnError(vars["key"], vars["value"])
//	encoder.setMessage("Value of " + vars["key"] + " changed")
//	encoder.encodeMessage(w)
//	log.Println(encoder.ActMessage)
//}
//
////re
//func HandlerDeleteRecord(w http.ResponseWriter, r *http.Request) {
//	var encoder Encoder
//	vars := mux.Vars(r)
//
//	err := strg.IsValueNotNullReturnErr(vars["key"])
//	if err != nil {
//		encoder.setError("Record doesn't have a value")
//		encoder.encodeError(w)
//		log.Println(encoder.ErrMessage)
//		return
//	}
//	strg.DeleteNullStorageRecords(vars["key"])
//	encoder.setMessage("Record deleted")
//	encoder.encodeMessage(w)
//	log.Println(encoder.ActMessage)
//}

func InitHandlersAndStartServe() {
	r := mux.NewRouter()
	//r.HandleFunc("/all", handlerShowAllRecords)
	r.HandleFunc("/{key}", handlerShowRecord)

	//r.HandleFunc("/setKey/{oldKey}/{newKey}", HandlerSetKey)
	//r.HandleFunc("/changeValue/{key}/{value}", HandlerChangeValue)
	//r.HandleFunc("/delete/{key}", HandlerDeleteRecord)
	//r.HandleFunc("/value/{key}", HandlerShowValue)

	log.Fatal(http.ListenAndServe(":8080", r))
}
