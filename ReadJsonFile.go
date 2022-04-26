package fxutils

import (
	"io/ioutil"
	"log"
)

// TODO BLOW AWAY

// should take in local file and go to *[]byte
// X func will call this to get whatever needed

// TODO remove old/misc notes!!

// I would like to think about combining all these files>>> Funcs into 1 file???

// Move to generic func
// remember w/ NerdCommenter is <leader>c<space> to toggle comments

// my <leader> is space bar (or key) whatever you want to refer to it as....so for me <space>"c"<space> is my toggle comment command

// func ReadJsonFile(f string) *[]byte {
//     data, err := ioutil.ReadFile(f)
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     return &data
// }
//
//thinking
// need to convert above func to generic func!

// remember !!!!!!
//

// going to convert on right to here!!!

// syntax should be :
// func FuncSig[x struct](p Params) *T <Pointer to struct> {}

// other examples i deref to struct...here I do not have struct...so deref to *[]byte

// THIS WORKS!!!!

//////////////////////////////////////////////////////////////////////////////

// ????? Generic *[]byte?????????
func ReadLocalJson[T any](p string) *[]byte {
	// param or -p- in this case should be a local file
	data, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	// deref to pointer where data is held
	// getting deref error b/c of type
	// long story short not pointer and not type
	return &data
}
