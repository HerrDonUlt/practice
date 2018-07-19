package handlers

import strg "practicegit/storage"
import "net/http"
import "github.com/gorilla/mux"
import "log"
import "encoding/json"

type JsonMes struct {
	Message string `json:"message"`
	Writer http.ResponseWriter
	Vars map[string]string
	KeyExisting bool	
	ValueExisting bool
	Error error
}

var mesArr JsonMes

func Init(){
	mesArr.KeyError = 
	mesArr.ValueError = 
}

func (jm JsonMes) wrap(w http.ResponseWriter, s string) {
	jm.Message = s
	jm.Writer = w
}

func (jm JsonMes) encode() {
	json.NewEncoder(jm.Writer).Encode(jm.Message)
}

func (jm JsonMes) log() {
	log.Println(jm.Message)
}

func (jm JsonMes) encodeMessage(w http.ResponseWriter, s string) {
	jm.wrap(w, s)
	jm.encode()
	jm.log()
}

func (jm JsonMes) getVars(r *http.Request) {
	jm.Vars = mux.Vars(r)
}

func (jm JsonMes) getExisting(s string) {
	jm.KeyExisting = strg.IsExist(jm.Vars[s])
}

func (jm JsonMes) sendFinalMessage(s string) {
	if jm.Existing {
		//when error
		jm.encodeMessage(jm.Writer, jm.Error)
	} else {
		jm.encodeMessage(jm.Writer, jm.Vars[s])
	}
}

func handlerShowRecord(w http.ResponseWriter, r *http.Request) {
	mesArr.getVars(r)
	mesArr.getKeyExisting("key")
	mesArr.sendFinalMessage("key")//w
}

func handlerReturnValue(w http.ResponseWriter, r *http.Request) {
	mesArr.getVars(r)
	mesArr.getValueExisting("value")
	mesArr.sendFinalMessage("value")
}

//re
func handlerKeySet(w http.ResponseWriter, r *http.Request) {
	mesArr.getVars(r)
	mesArr.getKeyExisting("oldKey")



	err := isKeyExist(oldKey)
	if err != "" {
		AddStorageRecord(oldKey, "")
		encodeAction(w, "New record set")
	} else {
		var oldValue *strg.KeyValInfo = Storage[oldKey]
		DeleteStorageRecord(oldKey)
		oldValue.SetKey(newKey)
		Storage[newKey] = oldValue
		encodeAction(w, "Key set")
	}

	AddLifetime(newKey)
}

//re
func handlerValueChange(w http.ResponseWriter, r *http.Request) {
	//need re
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	err := isKeyExist(key)
	if err != "" {
		AddStorageRecord(key, value)

		encodeAction(w, "New record set")
	} else {
		Storage[key].SetValue(value)
		AddLifetime(key)
		encodeAction(w, "Value set")
	}
}

//re
func handlerDeleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := isValueExist(vars["key"])
	if err != "" {
		encodeErr(w, err)
	} else {
		DeleteStorageRecord(vars["key"])
		encodeAction(w, "record with key:'"+vars["key"]+"' deleted")
	}
}

func HandleLoop() {
	r := mux.NewRouter()

	r.HandleFunc("/all", handlerShowAllJsonByKey)
	r.HandleFunc("/{key}", handlerShowJsonByKey)

	r.HandleFunc("/setkey/{oldKey}/{newKey}", handlerKeySet)
	r.HandleFunc("/changevalue/{key}/{value}", handlerValueChange)
	r.HandleFunc("/delete/{key}", handlerDeleteRecord)
	r.HandleFunc("/value/{key}", handlerReturnValue)

	log.Fatal(http.ListenAndServe(":8080", r))
}