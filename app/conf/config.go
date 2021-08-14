package conf

import (
	"cache/core"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	err := core.Injector.Provide(newConfig)
	core.CheckInjection(err, "newConfig")
}

type Config struct {
	Slots  int    `yaml:"slots"`
	TTL    int    `yaml:"ttl"`
	Policy string `yaml:"policy"`
}

func newConfig() *Config {
	file := os.Getenv("PATH_ENV")
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var c Config
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &c
}
