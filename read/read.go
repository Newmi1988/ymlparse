package read

import (
	"io/ioutil"
)



func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadYML(filePath string) []byte {
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	return dat
}
