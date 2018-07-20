package handlers

import strg "practicegit/storage"
import (
	"net/http"
	"log"
    "github.com/gorilla/mux"
)

//logging

func handlerShowRecord(w http.ResponseWriter, r *http.Request) {
	var encoder Encoder
	vars := mux.Vars(r)
	err := strg.IsKeyExistReturnErr(vars["key"])
	if err != nil {
		encoder.setError("Key is not exist")
		encoder.encodeError(w)
		log.Println(encoder.ErrMessage)
		return
	}
	encoder.setMessage("Record " + vars["key"] + " showed")
	encoder.encodeRecord(w, vars["key"])
	log.Println(encoder.ActMessage)
}

////re
//func handlerShowAllRecords(w http.ResponseWriter, r *http.Request) {
//	vars = mux.Vars(r)
//
//}

func handlerShowValue(w http.ResponseWriter, r *http.Request) {
	var encoder Encoder
	vars := mux.Vars(r)
	err := strg.IsValueNotNullReturnErr(vars["key"])
	if err != nil {
		encoder.setError("Value is not exist")
		encoder.encodeValue(w, vars["value"])
		log.Println(encoder.ErrMessage)
		return
	}
	encoder.setMessage("Value " + vars["key"] + " showed")
	encoder.encodeMessage(w)
	log.Println(encoder.ActMessage)
}

//re
func handlerSetKey(w http.ResponseWriter, r *http.Request) {
	var encoder Encoder
	vars := mux.Vars(r)
	err := strg.IsKeyExistReturnErr("oldKey")
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

func handlerChangeValue(w http.ResponseWriter, r *http.Request) {
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
func handlerDeleteRecord(w http.ResponseWriter, r *http.Request) {
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
