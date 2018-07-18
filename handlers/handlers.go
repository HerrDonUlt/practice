package handlers

import strt "practicegit/structs"
import "net/http"
import "github.com/gorilla/mux"
import "log"

func handlerShowAllJsonByKey(w http.ResponseWriter, r *http.Request) {
	encodeAllRecords(w)	
}

func handlerShowJsonByKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := isKeyExist(vars["key"])
	if err != "" {
		encodeErr(w, err)
	} else {
		encodeRecord(w, vars["key"])
	}
}

func handlerKeySet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	oldKey := vars["oldKey"]
	newKey := vars["newKey"]

	err := isKeyExist(oldKey)
	if err != "" {
		Storage[newKey].AddRecord(oldKey, "")
		encodeAction(w, "New record set")
	} else {
		var oldValue *strt.KeyValInfo = Storage[oldKey]
		delete(Storage, oldKey)
		oldValue.SetKey(newKey)
		Storage[newKey] = oldValue		
		encodeAction(w, "Key set")
	}

	extendLifetimeFn(Storage, newKey)
}

func handlerValueChange(w http.ResponseWriter, r *http.Request) {
	//need re
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	err := isKeyExist(key)
	if err != "" {		
		Storage[key].AddRecord(key, value)
		encodeAction(w, "New record set")
	} else {
		Storage[key].Value = value
		extendLifetimeFn(Storage, key)
		encodeAction(w, "Value set")		
	}
}

func handlerDeleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := isValueExist(vars["key"])
	if err != "" {
		encodeErr(w, err)
	} else {
		delete(Storage, vars["key"])
		encodeAction(w, "record with key:'" + vars["key"] + "' deleted")
	}
}

func Handle() {
	r := mux.NewRouter()

	r.HandleFunc("/all", handlerShowAllJsonByKey)
	r.HandleFunc("/{key}", handlerShowJsonByKey)

	r.HandleFunc("/setkey/{oldKey}/{newKey}", handlerKeySet)
	r.HandleFunc("/changevalue/{key}/{value}", handlerValueChange)
	r.HandleFunc("/delete/{key}", handlerDeleteRecord)

	log.Fatal(http.ListenAndServe(":8080", r))
}
