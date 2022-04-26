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

	// parse json
	o := parseJson[T](r)
	// TODO Remove bottom 3 comments
	// o := new(T)
	// json.Unmarshal(*r, o)
	//
	return o
}

// Read from local resoruce <json file>
func GetLocalJson[T any](p string) *T {
	// Try to get/read local file, else fail
	d, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	// parse json
	o := parseJson[T](&d)

	return o
}

// local only should return *[]
func parseJson[T any](b *[]byte) *T {
	o := new(T)
	json.Unmarshal(*b, o)

	return o
}
