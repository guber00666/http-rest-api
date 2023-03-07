package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	apiserver2 "http-rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver2.NewConfig()

	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	s := apiserver2.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
