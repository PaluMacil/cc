package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Config stores the available compilers and the selected default
type Config struct {
	Default   string
	Compilers []Compiler
}

// Load parses the Config from a file and returns it
func Load(path string) (*Config, error) {
	var c = new(Config)
	// Create empty config file if it hasn't been created yet
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = c.Save(path)
		if err != nil {
			return c, fmt.Errorf("initializing config: %s", err)
		}
	}
	// Read and parse config
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return &Config{}, fmt.Errorf("loading config: %s", err)
	}
	err = json.Unmarshal(raw, &c)
	if err != nil {
		return &Config{}, fmt.Errorf("parsing config: %s", err)
	}
	return c, nil
}

// Save stores the config in a file
func (c Config) Save(path string) error {
	b, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("saving config: %s", err)
	}
	ioutil.WriteFile(path, b, 0644)
	return nil
}

// SelectedCompiler returns whether there is a selected compiler
// and the compiler selected, if applicable.
func (c Config) SelectedCompiler() (bool, Compiler) {
	for _, compiler := range c.Compilers {
		if c.Default == compiler.Name {
			return true, compiler
		}
	}
	return false, Compiler{}
}

// Path returns the path to the config file by first checking the environment
// for CCUTIL_CONFIG_PATH and then checking the same folder as this binary.
func Path() string {
	configPath := os.Getenv("CCUTIL_CONFIG_PATH")
	if configPath == "" {
		configPath = filepath.Join(path.Dir(os.Args[0]), "config.json")
	}
	return configPath
}
