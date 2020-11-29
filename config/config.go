package config

import (
	"flag"
	"os"
)

type Config struct {
	Path string
	Debug bool
}

func Parse() *Config {

	c := Config{}

	path := flag.String("path", "", "Missing path parameter")
	debug := flag.Bool("debug", false, "Debug flag for verbose output")

	flag.Parse()

	if *path == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	c.Path = *path
	c.Debug = *debug

	return &c
}
