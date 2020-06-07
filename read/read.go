package read

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadYML(filePath string) string {
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	return string(dat)
}
