package storage

//-- Public methods below with mutexes
func (s *Storage) SubstructLifetimeOne(key string) {
	s.Lock()
	defer s.Unlock()
	s.substructLifeTimeOne(key)
}

func (s Storage) IsLifetimeZero(key string) bool {
	s.Lock()
	defer s.Unlock()
	return s.Record[key].isLifetimeZero()
}

func (s *Storage) Set(key, value string) {
	s.Lock()
	defer s.Unlock()
	s.Set(key, value)
	s.AddRecordLifetime(key)
}

func (s *Storage) AddRecordLifetime(key string) {
	s.Lock()
	defer s.Unlock()
	s.Record[key].addRecordLifetime()
}

func (s *Storage) GetRecord(key string) *Record {
	s.Lock()
	defer s.Unlock()
	return s.getRecord(key)
}

func (s *Storage) GetRecordValue(key string) string {
	s.Lock()
	defer s.Unlock()
	return s.getRecord(key).Value
}

func (s *Storage) Count() int {
	s.Lock()
	defer s.Unlock()
	return s.count()
}

func (s *Storage) IsKeyInStorage(key string) bool {
	s.Lock()
	defer s.Unlock()
	return s.isKeyInStorage(key)
}

func (s *Storage) IsValueInRecordNotNull(key string) bool {
	s.Lock()
	defer s.Unlock()
	return s.isValueInRecordNotNull(key)
}
