package structs

import "time"
import "fmt"

const sleeptimeInSec time.Duration = 4 * time.Second
const stdlifetime int = 3

//store the info about key-val thing
type KeyValInfo struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	LifeTime int    `json:"life_time"`
}

func (kv *KeyValInfo) SubstructOne() {
	fmt.Println(kv.LifeTime)
	kv.LifeTime = kv.LifeTime - 1
}

func (kv KeyValInfo) IsZero() bool {
	if kv.LifeTime == 0{
		return true
	} else {
		return false
	}
}

func (kv *KeyValInfo) extendLifetimeFn(key string) {
	kv.LifeTime += stdlifetime
}

func (kv *KeyValInfo) isKeyExist(k string) bool {
	if k == kv.Key {
		return true
	}
	
	return false
}

func isKeyExist(k string) bool {
	for _, item := range Storage {
		item.isKeyExist()
	}
	return false
}

func (kv KeyValInfo) Delete(m map[string]*KeyValInfo, k string) {
	delete(m, k)
}

func LifetimeManage(storage map[string]*KeyValInfo) {
	for {
		time.Sleep(sleeptimeInSec)
		for _, s := range storage {
			isZero := s.IsZero()
			fmt.Println(isZero)
			if isZero  {
				s.Delete(storage, s.Key)

			} else {
				s.SubstructOne()
			}
		}
	}
}
