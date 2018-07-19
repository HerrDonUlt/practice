package storage

import "sync"

var storage = make(map[string]*Storage)

type Record struct {
	sync.Mutex
	Key      string `json:"key"`
	Value    string `json:"value"`
}

//store the info about key-val thing
type Storage struct {
	sync.Mutex
	Record    Record
	LifeTime int    `json:"life_time"`
}

func (s *Storage) substructLifetimeOne() {
	s.Lock()
	defer s.Unlock()
	s.LifeTime -= 1
}

func (s Storage) isLifetimeZero() bool {
	s.Lock()
	defer s.Unlock()
	if s.LifeTime == 0 {
		return true
	} 
	return false
}

func (p Record) isKeyIn(k string) bool {
	p.Lock()
	defer p.Unlock()
	if p.Key == k {
		return true
	}
	return false
}

func (p Record) isValueIn(k string) bool {
	p.Lock()
	defer p.Unlock()
	if p.Key == k && p.Value != "" {
		return true
	}
	return false
}

func (p *Record) setKey(k string) {
	p.Lock()
	defer p.Unlock()
	p.Key = k
}

func (p *Record) setValue(v string) {
	p.Lock()
	defer p.Unlock()
	p.Value = v
}

func (s *Storage) addRecordLifetime() {
	s.Lock()
	defer s.Unlock()
	s.LifeTime += stdlifetime
}
