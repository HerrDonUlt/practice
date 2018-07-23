package storage

import (
	"time"
)

const sleeptimeInSec time.Duration = 4 * time.Second

func InitTestStorage() {
	SetRecord("1", "something")
	SetRecord("2", "something new")
	SetRecord("3", "")
}

func LifetimeManage() {
	for {
		time.Sleep(sleeptimeInSec)
		DeleteNullStorageRecords()
		SubstructLifetimeRecords()
	}
}
