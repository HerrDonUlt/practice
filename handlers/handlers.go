package handlers

import strg "practicegit/storage"
import "net/http"
import "github.com/gorilla/mux"
import "log"
import "encoding/json"
import "errors"

type ErrMsg struct {
	Message error `json:"error_message"`
}
type ActMsg struct {
	Message string `json:"action_message"`
}

type JsonMes struct {
	ErrMessage ErrMsg
	ActMessage ActMsg
	Vars map[string]string
	Existing bool
	Error error
}

var mesArr JsonMes

func (jm *JsonMes) setErr(s string) {
	jm.ErrMessage.Message = errors.New(s)
}

func (jm *JsonMes) setMes(s string) {
	jm.ActMessage.Message = s
}

func (jm JsonMes) encodeErr(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(jm.ErrMessage)
}

func (jm JsonMes) encodeRecord(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(strg.ReturnStorageRecord(jm.Vars["key"]))
}

func (jm JsonMes) encodeAllRecords(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(strg.ReturnStorage())
}

func (jm JsonMes) encodeValue(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(strg.ReturnValueRecord(jm.Vars["key"]))
}

func (jm JsonMes) encodeMes(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(jm.ActMessage)
}

func (jm JsonMes) logErr() {
	log.Println(jm.ErrMessage)
}

func (jm JsonMes) logMes() {
	log.Println(jm.ActMessage)
}

func (jm *JsonMes) getVars(r *http.Request) {
	jm.Vars = mux.Vars(r)
}

func (jm JsonMes) getExisting(s string) {
	jm.Existing, jm.Error = strg.IsExist(s, jm.Vars["key"])
}

func HandleLoop() {
	r := mux.NewRouter()

	r.HandleFunc("/all", handlerShowAllRecords)
	r.HandleFunc("/{key}", handlerShowRecord)

	r.HandleFunc("/setkey/{oldKey}/{newKey}", handlerKeySet)
	r.HandleFunc("/changevalue/{key}/{value}", handlerValueChange)
	r.HandleFunc("/delete/{key}", handlerDeleteRecord)
	r.HandleFunc("/value/{key}", handlerShowValue)

	log.Fatal(http.ListenAndServe(":8080", r))
}