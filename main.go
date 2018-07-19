//golang race
//postman

//нормальные названия файлов
//упростить isExist разбить
//убрать глобальную перменную месарр(переименовать)
//нормальные ветвления
//переделать ретернсторадж
//убрать ненужные функции в джсонмес
//otdelit loggirovanie
//peredelat storage

package main

import (
	hdlr "practicegit/handlers"
	strg "practicegit/storage"
	// "fmt"
)

func main() {
	strg.TestStorageAdding()
	go strg.LifetimeManage()
	hdlr.HandleEstalish()
}
