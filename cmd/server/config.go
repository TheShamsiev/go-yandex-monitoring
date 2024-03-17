package main

import (
	"flag"
	"os"
)

type Config struct {
	Address string
}

func ParseConfig() Config {
	var address string

	flag.StringVar(&address, "a", "localhost:8080", "address with port for the server to run on")
	flag.Parse()

	if envAddress := os.Getenv("ADDRESS"); envAddress != "" {
		address = envAddress
	}

	return Config{address}
}
