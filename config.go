package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"sigs.k8s.io/yaml"
)

type config struct {
	path   string
	cfg    interface{}
	envVar *string
	yaml   bool
}

// New reads a config file and populates the supplied config object
func New(cfg interface{}, path string, options ...func(*config) error) error {
	if cfg == nil {
		return fmt.Errorf("supplied config object is nil")
	}
	c := new(config)
	c.path = path
	c.cfg = cfg

	for _, opt := range options {
		opt(c)
	}

	if c.envVar != nil {
		c.path = os.Getenv(*c.envVar)
	}

	if len(c.path) == 0 {
		return fmt.Errorf("empty config path")
	}

	return c.get()
}

func (c config) get() error {
	b, err := ioutil.ReadFile(c.path)
	if err != nil {
		return err
	}

	if c.yaml {
		return yaml.Unmarshal(b, c.cfg)
	}
	return json.Unmarshal(b, c.cfg)
}

// Yaml indicates that the config file is in YAML format
func Yaml(c *config) error {
	c.yaml = true
	return nil
}

// EnvVar if set will override the config path
func EnvVar(name string) func(*config) error {
	return func(c *config) error {
		if len(name) == 0 {
			return fmt.Errorf("empty envvar name")
		}
		c.envVar = &name
		return nil
	}
}
