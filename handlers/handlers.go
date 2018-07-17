 package handlers

import s "practicegit/structs"
import "net/http"
import "encoding/json"
import "github.com/gorilla/mux"
import "log"

const stdlifetime int = 3

var Storage map[string]*s.KeyValInfo

type JsonErrMes struct {
	ErrMessage string `json:"err_message`
}

func (jem JsonErrMes) setErrMes(s string) {
	jem.ErrMessage = s
}

func (jem JsonErrMes) encodeErrMes(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(jem.ErrMessage)
}

func (jem JsonErrMes) logErr(w http.ResponseWriter, k string) {
	jem.setErrMes(k)
	jem.endcodeErrMes(w)
}

//don't erase
// s.setErrMes("New Key is not unique")
// 			s.endcodeErrMes(w)

func handlerShowAllJsonByKey(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Storage)
}

func handlerShowJsonByKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	for _, item := range Storage {
		item.isKeyExist()
	}
	isExist := isKeyExist(key)

	if !isExist {
		panic("this key doesn't exist")
	}

	json.NewEncoder(w).Encode(Storage[key])
}

func handlerKeySet(w http.ResponseWriter, r *http.Request) {
	//dont foget to check newKey unique
	vars := mux.Vars(r)
	oldKey := vars["oldKey"]
	newKey := vars["newKey"]

	isNewKeyUnique(newKey)
	if isKeyExist(oldKey) {
		var oldValue *s.KeyValInfo = Storage[oldKey]
		delete(Storage, oldKey)
		oldValue.Key = newKey
		Storage[newKey] = oldValue
	} else {
			Storage[oldKey] = &s.KeyValInfo{oldKey,"", stdlifetime}
	}

	extendLifetimeFn(newKey)
}

func handlerValueChange(w http.ResponseWriter, r *http.Request) {	
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]
	if isKeyExist(key) {
		Storage[key].Value = value
	 	extendLifetimeFn(key)
	} else {
	 	Storage[key] = &s.KeyValInfo{key,value,stdlifetime}
	}
}

func handlerDeleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	if Storage[key].Value == "" {
		panic("Record doesn't have a value")
	} else{
		delete(Storage, key)	
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
