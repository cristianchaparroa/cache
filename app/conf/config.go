package conf

import (
	"cache/core"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	err := core.Injector.Provide(newConfig)
	core.CheckInjection(err, "newConfig")
}

const (
	OlderFistEvictionPolicy   = "OLDEST_FIRST"
	NewestFirstEvictionPolicy = "NEWEST_FIRST"
	RejectEvictionPolicy      = "REJECT"
)

type Config struct {
	Slots  int    `yaml:"slots"`
	TTL    int64  `yaml:"ttl"`
	Policy string `yaml:"policy"`
}

func newConfig() (*Config, error) {
	file := os.Getenv("CONFIG")
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return nil, nil
	}

	var c Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}

	err = c.IsValid()

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Config) IsValid() error {
	if c.Slots <= 0 {
		return errors.New("the number of slots is negative")
	}

	if c.TTL <= 0 {
		return errors.New("the TTL value is negative")
	}

	validPolicies := []string{OlderFistEvictionPolicy, NewestFirstEvictionPolicy, RejectEvictionPolicy}

	isValidPolicy := false
	for _, policy := range validPolicies {
		if policy == c.Policy {
			isValidPolicy = true
		}
	}

	if !isValidPolicy {
		return errors.New("policy not expected")
	}

	return nil
}
