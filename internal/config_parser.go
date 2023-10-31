package internal

import (
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func ParseConfig() {
	var config Config
	data := []byte(`
database:
  user: test
  password: test
`)
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Value: %#v\n", config)
}
