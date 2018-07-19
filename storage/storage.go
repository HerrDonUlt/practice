
package storage



var storage = make(map[string]*Storage)

//store the info about key-val thing
type Storage struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	LifeTime int    `json:"life_time"`
}

func (s *Storage) substructOne() {
	s.LifeTime -= 1
}

func (s Storage) isZero() bool {
	if s.LifeTime == 0 {
		return true
	} else {
		return false
	}
}

func (s Storage) isKeyIn(k string) bool {
	if s.Key == k {
		return true
	} else {
		return false
	}
}

func (s Storage) isValueIn(k string) bool {
	if s.Key == k && s.Value != "" {
		return true
	} else {
		return false
	}
}

func (s Storage) delete(k string) {
	delete(storage, k)
}

func (s *Storage) setKey(k string) {
	s.Key = k
}

func (s *Storage) setValue(v string) {
	s.Value = v
}

func (s *Storage) addStorageRecord(k string, v string) {
	storage[k] = &Storage{k, v, stdlifetime}
}

func (s *Storage) addLifetime(k string) {
	storage[k].LifeTime += stdlifetime
}