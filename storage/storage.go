package storage



var storage = make(map[string]*Storage)

//store the info about key-val thing
type Storage struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	LifeTime int    `json:"life_time"`
}

func (kv *Storage) SubstructOne() {
	kv.LifeTime -= 1
}

func (kv Storage) IsZero() bool {
	if kv.LifeTime == 0 {
		return true
	} else {
		return false
	}
}

func (kv Storage) IsKeyIn(k string) bool {
	if kv.Key == k {
		return true
	} else {
		return false
	}
}

func (kv Storage) IsValueIn(k string) bool {
	if kv.Key == k && kv.Value != "" {
		return true
	} else {
		return false
	}
}

func (kv Storage) Delete(k string) {
	delete(storage, k)
}

func (kv *Storage) SetKey(k string) {
	kv.Key = k
}

func (kv *Storage) SetValue(v string) {
	kv.Value = v
}

func (kv *Storage) AddStorageRecord(k string, v string) {
	storage[k] = &Storage{k, v, stdlifetime}
}

func (kv *Storage) AddLifetime(k string) {
	storage[k].LifeTime += stdlifetime
}
