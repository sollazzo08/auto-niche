package config

import (
	"log"
)

type ConfigSettings struct {

}

func LoadConfig() (*ConfigSettings, error){

	log.Fatal("Hello World")
	return &ConfigSettings {

	}, nil
}
