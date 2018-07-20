package storage

const stdlifetime int = 3

type Record struct {
	Key string `json:"key"`
	Value string `json:"value"`
	LifeTime int    `json:"life_time"`
}

func (r *Record) substructLifetimeOne() {
	r.LifeTime -= 1
}

func (r Record) isLifetimeZero() bool {
	if r.LifeTime == 0 {
		return true
	}
	return false
}

func (r *Record) set(key, value string) {
	r.Key = key
	r.Value = value
}

func (r *Record) addRecordLifetime() {
	r.LifeTime += stdlifetime
}

func (r *Record) getRecordKey() string {
	return r.Key
}

func (r *Record) getRecordValue() string {
	return r.Value
}
