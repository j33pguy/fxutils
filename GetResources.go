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

	o := parseJSON[T](r)

	return o
}

// Read from local resoruce <json file>
func GetLocalJSON[T any](p string) *T {
	// Try to get/read local file, else fail
	d, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	o := parseJSON[T](&d)

	return o
}

// local only....should return *T
func parseJSON[T any](b *[]byte) *T {
	o := new(T)
	json.Unmarshal(*b, o)

	return o
}
