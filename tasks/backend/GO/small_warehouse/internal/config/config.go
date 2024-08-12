package config

import "os"

type Config struct {
	HttpAddr string
}

func ReadEnv() *Config {
	conf := &Config{}
	httpAddr, ok := os.LookupEnv("HTTP_ADDR")
	if ok {
		conf.HttpAddr = httpAddr
	} else {
		conf.HttpAddr = "0.0.0.0:3000"
	}

	return conf
}
