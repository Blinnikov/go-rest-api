package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/blinnikov/go-rest-api/internal/app/apiserver"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	log.Println("Main program start")
	flag.Parse()

	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := getLogger(config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	logger.Println("Logger configured")

	if err := apiserver.Start(config, logger); err != nil {
		log.Fatal(err)
	}
}

func getConfig() (*apiserver.Config, error) {
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		return nil, err
	}

	return config, err
}

func getLogger(logLevel string) (*logrus.Logger, error) {
	logger := logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(level)
	return logger, nil
}
