package structs

import "time"

const sleeptimeInSec time.Duration = 4 * time.Second
const stdlifetime int = 3

//store the info about key-val thing
type KeyValInfo struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	LifeTime int    `json:"life_time"`
}

func (kv *KeyValInfo) SubstructOne() {
	kv.LifeTime -= 1
}

func (kv KeyValInfo) IsZero() bool {
	if kv.LifeTime == 0 {
		return true
	} else {
		return false
	}
}

func (kv KeyValInfo) IsKeyIn(k string) bool {
	if kv.Key == k {
		return true
	} else {
		return false
	}	
}

func (kv KeyValInfo) IsValueIn(k string) bool {
	if kv.Key == k && kv.Value != "" {
		return true
	} else {
		return false
	}	
}

func (kv KeyValInfo) Delete(m map[string]*KeyValInfo, k string) {
	delete(m, k)
}

func (kv KeyValInfo) SetKey(k string) {
	kv.Key = k
}

func LifetimeManage(storage map[string]*KeyValInfo) {
	for {
		time.Sleep(sleeptimeInSec)
		for _, s := range storage {
			if s.IsZero() {
				s.Delete(storage, s.Key)
			} else {
				s.SubstructOne()
			}
		}
	}
}
