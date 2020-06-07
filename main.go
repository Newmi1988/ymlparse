package main


import "fmt"
import "github.com/Newmi1988/ymlparse/read"
import "github.com/Newmi1988/ymlparse/config"

func main() {
	fmt.Println("Config Parser")
	cfg := config.Config{}

	data := read.ReadYML("test.yml")

	cfg.SetFromBytes(data)
}


