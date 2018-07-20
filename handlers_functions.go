package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerShowRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	json.NewEncoder(w).Encode(vars["key"])
	log.Println("Record " + vars["key"] + " showed")
}

////re
//func handlerShowAllRecords(w http.ResponseWriter, r *http.Request) {
//	vars = mux.Vars(r)
//
//}

func HandlerShowValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	json.NewEncoder(w).Encode(vars["value"])
	log.Println("Value " + vars["key"] + " showed")
}

//re
func HandlerSetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if strg.IsKeyInStorage(vars["key"]) {

	}
	if err != nil {
		strg.AddStorageRecord(vars["oldKey"], "")
		encoder.setMessage("New record seted")
		encoder.encodeMessage(w)
		log.Println(encoder.ActMessage)
		return
	}
	err = strg.IsKeyExistReturnErr("newKey")
	if err != nil {
		encoder.setError("New key already in use")
		encoder.encodeValue(w, vars["key"])
		log.Println(encoder.ErrMessage)
		return
	}
	strg.ChangeRecordKeyReturnErr(vars["oldKey"], vars["newKey"])
	encoder.setMessage("Key changed")
	encoder.encodeMessage(w)
	log.Println(encoder.ActMessage)
}

func HandlerChangeValue(w http.ResponseWriter, r *http.Request) {
	var encoder Encoder
	vars := mux.Vars(r)

	err := strg.IsKeyExistReturnErr(vars["key"])
	if err != nil {
		encoder.setError("Key is not exist")
		encoder.encodeError(w)
		log.Println(encoder.ErrMessage)
		return
	}
	strg.ChangeRecordValueReturnError(vars["key"], vars["value"])
	encoder.setMessage("Value of " + vars["key"] + " changed")
	encoder.encodeMessage(w)
	log.Println(encoder.ActMessage)
}

//re
func HandlerDeleteRecord(w http.ResponseWriter, r *http.Request) {
	var encoder Encoder
	vars := mux.Vars(r)

	err := strg.IsValueNotNullReturnErr(vars["key"])
	if err != nil {
		encoder.setError("Record doesn't have a value")
		encoder.encodeError(w)
		log.Println(encoder.ErrMessage)
		return
	}
	strg.DeleteStorageRecord(vars["key"])
	encoder.setMessage("Record deleted")
	encoder.encodeMessage(w)
	log.Println(encoder.ActMessage)
}
