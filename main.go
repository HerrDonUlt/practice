//split to packeges

package main

import (
	"encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	//"json"
	//"net/url"
	// "html/template"
	// "io/ioutil"
)

//default lifetime value add to element o storage
const stdlifetime int = 3
const stdsleeptime time.Duration = 4

// var extendLifetimeChannel chan string = make (chan string)

//store the info about key-val thing
type KeyValInfo struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	LifeTime int    `json:"life_time"`
}

var storage = make(map[string]*KeyValInfo)

//substract the lifetime and delete zero lifetime KeyValInfo
func lifetimeManage(storage map[string]*KeyValInfo) {
	for {
		time.Sleep(stdsleeptime * time.Second)
		for _, s := range storage {
			if s.LifeTime == 0 {
				delete(storage, s.Key)

			} else {
				s.LifeTime -= 1
			}
		}
	}
}

func testStorageAdding() {
	storage["1"] = &KeyValInfo{"1", "something", stdlifetime}
	storage["2"] = &KeyValInfo{"2", "something new", stdlifetime + 2}
	storage["3"] = &KeyValInfo{"3", "something new too", stdlifetime}
}

func extnedLifetimeFn(key string) {
	storage[key].LifeTime += stdlifetime
}

func checkNewKeyUnique(k string) {
	for _, s := range storage {
		if k == s.Key {
			panic("New Key is not unique")
		}
	}
}

func handlerShowAllJsonByKey(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage)
}

func handlerShowJsonByKey(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["key"]

	json.NewEncoder(w).Encode(storage[key])
}

func handlerKeyChange(w http.ResponseWriter, r *http.Request) {
	//dont foget to check newKey unique
	vars := mux.Vars(r)
	oldKey := vars["oldKey"]
	newKey := vars["newKey"]

	checkNewKeyUnique(newKey)

	var oldValue *KeyValInfo = storage[oldKey]
	delete(storage, oldKey)
	oldValue.Key = newKey
	storage[newKey] = oldValue

	extnedLifetimeFn(newKey)
}

func handlerValueChange(w http.ResponseWriter, r *http.Request) {
	//dont foget to check key exist
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	storage[key].Value = value

	extnedLifetimeFn(key)
}

func main() {

	testStorageAdding()

	r := mux.NewRouter()

	r.HandleFunc("/all", handlerShowAllJsonByKey)
	r.HandleFunc("/{key}", handlerShowJsonByKey)
	r.HandleFunc("/changekey/{oldKey}/{newKey}", handlerKeyChange)
	r.HandleFunc("/changevalue/{key}/{value}", handlerValueChange)

	go lifetimeManage(storage)

	log.Fatal(http.ListenAndServe(":8080", r))
}
