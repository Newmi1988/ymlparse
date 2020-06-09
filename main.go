package main


import (
	"fmt"
	"log"

	"github.com/Newmi1988/ymlparse/read"
	"github.com/Newmi1988/ymlparse/config"
)

func main() {
	fmt.Println("Config Parser")
	cfg := config.Config{}

	data := read.ReadYML("test.yml")

	cfg.SetFromBytes(data)

	serviceConfig, err := cfg.Get("key2")
	if err != nil {
		log.Fatalf("Error with config")
	}

	fmt.Println(serviceConfig)
}


