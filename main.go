//1 copy the google golang web-server example !done
//2 create map element !done
//2.1 map with structure !done
//2.1 how to get data - by the url
//2.3 q: how to name a key
//3 add a life time to key
//3.1  extend key lifetime when value is change
//3.2  delete key when lifetime is end
//4 add a fn that set a lifetime to key
//5 add a fn that return a value by key
//6 handle err of deleting null value key

//---------------
//timer for lifetime

//---------------
//how to detect a change - 

//gorilla mux

package main

import (
	"fmt"
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

//store the info about key-val thing
type KeyValInfo struct {
	Key      string
	Value    string
	LifeTime int
}

var storage = make(map[string]*KeyValInfo)

//substract the lifetime and delete zero lifetime KeyValInfo
func lifetimeManage(storage map[string]*KeyValInfo) {
	for {
		time.Sleep(4 * time.Second)
		for _, s := range storage {
			if s.LifeTime == 0 {
				delete(storage, ""+s.Key)

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

func handler(w http.ResponseWriter, r *http.Request) {

	for _, s := range storage {
		fmt.Fprintf(w, "%s %s %d\n", s.Key, s.Value, s.LifeTime)
	}
}

func main() {

	testStorageAdding()

	go lifetimeManage(storage)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
