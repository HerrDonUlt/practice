package handlers

import strg "practicegit/storage"
import "net/http"

func handlerShowRecord(w http.ResponseWriter, r *http.Request) {
	mesArr.getVars(r)
	mesArr.getExisting("key")
	if mesArr.Existing {
		mesArr.setErr("Key is not exist")
		mesArr.encodeErr(w)
		mesArr.logErr()
	} else {
		mesArr.setMes("Record " + mesArr.Vars["key"] + " showed")
		mesArr.encodeRecord(w)
		mesArr.logMes()
	}
}

//re
func handlerShowAllRecords(w http.ResponseWriter, r *http.Request) {
	mesArr.getVars(r)
	mesArr.getExisting("key")
	if mesArr.Existing {
		mesArr.setErr("Key is not exist")
		mesArr.encodeErr(w)
		mesArr.logErr()
	} else {
		mesArr.setMes("All records showed")
		mesArr.encodeAllRecords(w)
		mesArr.logMes()
	}
}

func handlerShowValue(w http.ResponseWriter, r *http.Request) {
	mesArr.getVars(r)
	mesArr.getExisting("value")
	if mesArr.Existing {
		mesArr.setErr("Value is not exist")
		mesArr.encodeErr(w)
		mesArr.logErr()
	} else {
		mesArr.setMes("Value " + mesArr.Vars["key"] + " showed")
		mesArr.encodeValue(w)
		mesArr.logMes()
	}
}

//re
func handlerKeySet(w http.ResponseWriter, r *http.Request) {
	mesArr.getVars(r)
	mesArr.getExisting("oldKey")
	if mesArr.Existing {
		strg.AddStorageRecord(mesArr.Vars["oldKey"], "")
		mesArr.setMes("New record seted")
		mesArr.encodeMes(w)
		mesArr.logMes()
	} else {
		strg.ChangeStorageKey(mesArr.Vars["oldKey"], mesArr.Vars["newKey"])		
		mesArr.setMes("Key changed")
		mesArr.encodeMes(w)
		mesArr.logMes()
	}
}

func handlerValueChange(w http.ResponseWriter, r *http.Request) {
	mesArr.getVars(r)
	mesArr.getExisting("key")
	if mesArr.Existing {
		mesArr.setErr("Record doesn't have a value")
		mesArr.encodeErr(w)
		mesArr.logErr()
	} else {
		strg.ChangeStorageValue(mesArr.Vars["key"], mesArr.Vars["value"])
		mesArr.setMes("Value of " + mesArr.Vars["key"] + " changed")
		mesArr.encodeMes(w)
		mesArr.logMes()
	}
}

//re
func handlerDeleteRecord(w http.ResponseWriter, r *http.Request) {
	mesArr.getVars(r)
	mesArr.getExisting("value")
	if mesArr.Existing {
		mesArr.setErr("Record doesn't have a value")
		mesArr.encodeErr(w)
		mesArr.logErr()
	} else {
		strg.DeleteStorageRecord(mesArr.Vars["key"])
		mesArr.setMes("Record deleted")
		mesArr.encodeMes(w)
		mesArr.logMes()
	}
}