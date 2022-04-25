package fxutils

import (
	"io/ioutil"
	"log"
)

func ReadJsonFile(f string) *[]byte {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}

	return &data
}
