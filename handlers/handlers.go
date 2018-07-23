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
	record := storage.GetRecord(vars["key"])
	json.NewEncoder(w).Encode(record)
	log.Println("Record " + vars["key"] + " showed")
}

//func handlerShowAllRecords(w http.ResponseWriter, r *http.Request) {
//	records := storage.Strg.GetAllRecord()
//	for _, s := range records {
//
//		json.NewEncoder(w).Encode(s)
//
//	}
//	log.Println("showed all records")
//}

func HandlerSetValueForRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]
	if storage.IsKeyInStorage(key) {
		storage.SetRecord(key, value)
		storage.AddLifetimeForRecord(key)
		json.NewEncoder(w).Encode("Record change value: " + value + " with key: " + key)
		log.Println("Record change value: " + value + " with key: " + key)
		return
	}
	storage.SetRecord(key, value)
	json.NewEncoder(w).Encode("New record set with key: " + key)
	log.Println("New record set with key: " + key)
}

func HandlerReturnValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	if storage.IsKeyInStorage(key) {
		if storage.IsValueInStorageNotNull(key) {
			value := storage.GetRecordValue(storage.GetRecord(key))
			storage.AddLifetimeForRecord(key)
			json.NewEncoder(w).Encode(value)
			log.Println("Value: " + value + "showed with key: " + key)
			return
		}
		json.NewEncoder(w).Encode("Value is not exist")
		log.Println("Value is not exist")
		return
	}
	json.NewEncoder(w).Encode("Key is not exist")
	log.Println("Key is not exist")
	return
}

//re
func HandlerDeleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	if storage.IsKeyInStorage(key) {
		if storage.IsValueInStorageNotNull(key) {
			storage.DeleteStorageRecord(key)
			json.NewEncoder(w).Encode("Record deleted with key: " + key)
			log.Println("Record deleted with key: " + key)
			return
		}
		json.NewEncoder(w).Encode("Value is not exist")
		log.Println("Value is not exist")
		return
	}
	json.NewEncoder(w).Encode("Key is not exist")
	log.Println("Key is not exist")
	return
}

func InitHandlersAndStartServe() {
	r := mux.NewRouter()
	//r.HandleFunc("/all", handlerShowAllRecords)
	r.HandleFunc("/{key}", handlerShowRecord)

	r.HandleFunc("/setValue/{key}/{value}", HandlerSetValueForRecord)
	r.HandleFunc("/changeValue/{key}/{value}", HandlerReturnValue)
	r.HandleFunc("/delete/{key}", HandlerDeleteRecord)

	log.Fatal(http.ListenAndServe(":8080", r))
}
