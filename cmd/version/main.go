package main

import (
	"log"

	"github.com/unixoff/goparse/internal/version"
)

func main() {
	conf, err := version.NewConfig()
	if err != nil {
		log.Fatalln("Config error:", err)
	}

	collect := version.NewCollect(conf)
	collect.Run()
}
