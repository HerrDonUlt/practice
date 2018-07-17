package structs

//store the info about key-val thing
type KeyValInfo struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	LifeTime int    `json:"life_time"`
}