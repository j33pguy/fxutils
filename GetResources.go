package fxutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

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
	data, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	return &data
}
