package fxutils

import (
	"io/ioutil"
	"log"
	"net/http"
)

// generic not needed

func Response(baseurl string) *[]byte {
	response, err := http.Get(baseurl)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return &responseData
}
