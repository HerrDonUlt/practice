package handlers

const stdlifetime int = 3

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