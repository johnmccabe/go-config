package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	IntKey       int
	StringKey    string
	BoolKey      bool
	StringArrKey []string
	MapKey       map[string]string
}

var testConfigValid = TestConfig{
	IntKey:    1234,
	StringKey: "stringval",
	BoolKey:   true,
	StringArrKey: []string{
		"arrVal1",
		"arrVal2",
	},
	MapKey: map[string]string{
		"subKey": "subkeyval",
	},
}

func TestJSONValid(t *testing.T) {
	cfg := new(TestConfig)
	New(cfg, "testdata/config.json")
	assert.Equal(t, testConfigValid, *cfg, "config doesn't match expected values")
}

func TestYAMLValid(t *testing.T) {
	cfg := new(TestConfig)
	New(cfg, "testdata/config.yaml", Yaml)
	assert.Equal(t, testConfigValid, *cfg, "config doesn't match expected values")
}

func TestEnvVarValid(t *testing.T) {
	cfg := new(TestConfig)
	os.Setenv("CONFIG_OVERRIDE", "testdata/config.json")
	New(cfg, "testdata/config.doesnotexist", EnvVar("CONFIG_OVERRIDE"))
	assert.Equal(t, testConfigValid, *cfg, "config doesn't match expected values")
}
