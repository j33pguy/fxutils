package fxutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// DONE checking to make sure i spelled resources correctly <<<< SMOOTH BRAIN!!!

// TODO recheck all spelling of files after today.....too many booze<did that on purpose>

// call this file ReadResource???
// TODO think about naming for clarity

// Read from foxholeservices warapi
// FuncSig[Struct](args/params) Return<struct> {}
func GetFXApi[T any](p string) *T {
	b := Baseurl(p)
	r := Response(*b)

	o := new(T)
	json.Unmarshal(*r, o)

	return o
}

// Read from local resoruce <json file>
// FuncSig[Struct](args/params) Return<*[]byte> {}
func GetLocalJson[T any](p string) *[]byte {
	// mv data,err to package??? is it repeating???
	data, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	return &data

}
