package fxutils

import (
	"encoding/json"
)

func GetFuncResourceApi[T any](p string) *T {
	b := Baseurl(p)
	r := Response(*b)

	o := new(T)
	json.Unmarshal(*r, o)

	return o
}
