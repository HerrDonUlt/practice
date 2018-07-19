package handlers

import strg "practicegit/storage"
import "net/http"
import "github.com/gorilla/mux"
import "log"
import "encoding/json"
import "errors"

type Encoder struct {
	ErrMessage error `json:"error_message"`
	ActMessage string `json:"action_message"`
}

func (e *Encoder) setError(s string) {
	e.ErrMessage = errors.New(s)
}

func (e *Encoder) setMessage(s string) {
	e.ActMessage = s
}

func (e Encoder) encodeError(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(e.ErrMessage)
}

func (e Encoder) encodeRecord(w http.ResponseWriter, key string) {
	json.NewEncoder(w).Encode(strg.ReturnStorageRecord(key))
}

func (e Encoder) encodeValue(w http.ResponseWriter, value string) {
	json.NewEncoder(w).Encode(strg.ReturnRecordValue(value))
}

func (e Encoder) encodeMessage(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(e.ActMessage)
}

func HandleEstalish() {
	r := mux.NewRouter()

	//r.HandleFunc("/all", handlerShowAllRecords)
	r.HandleFunc("/{key}", handlerShowRecord)

	r.HandleFunc("/setkey/{oldKey}/{newKey}", handlerKeySet)
	r.HandleFunc("/changevalue/{key}/{value}", handlerValueChange)
	r.HandleFunc("/delete/{key}", handlerDeleteRecord)
	r.HandleFunc("/value/{key}", handlerShowValue)

	log.Fatal(http.ListenAndServe(":8080", r))
}
