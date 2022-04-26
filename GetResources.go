package fxutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// TODO recheck all spelling of files after today.....too many booze<did that on purpose>
// TODO think about naming for clarity

// Read from foxholeservices warapi
func GetFXApi[T any](p string) *T {
	b := Baseurl(p)
	r := Response(*b)

	o := new(T)
	json.Unmarshal(*r, o)

	return o
}

// Read from local resoruce <json file>
func GetLocalJson[T any](p string) *[]byte {
	// mv data,err to package??? is it repeating???
	data, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	return &data

}
