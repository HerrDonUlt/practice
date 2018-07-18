package handlers

import "encoding/json"
import "log"
import "net/http"

type JsonErrMes struct {
	ErrMessage string `json:"err_message"`
}

type JsonActMes struct {
	ActionMessage string `json:"action_message"`
}

func encodeAllRecords(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(Storage)
	log.Println("show all records")
}

func encodeRecord(w http.ResponseWriter, s string) {
	json.NewEncoder(w).Encode(Storage[s])
	log.Println("show record with key:'" + s + "'")
}

func encodeValue(w http.ResponseWriter, s string) {
	json.NewEncoder(w).Encode(Storage[s].Value)
	log.Println("Return value by key'" + s + "'")
}

func encodeErr(w http.ResponseWriter, s string) {
	json.NewEncoder(w).Encode(JsonErrMes{s})
	log.Println(s)
}

func encodeAction(w http.ResponseWriter, s string) {
	json.NewEncoder(w).Encode(JsonActMes{s})
	log.Println(s)
}
