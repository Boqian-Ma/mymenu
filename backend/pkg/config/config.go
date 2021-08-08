package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// Default config values
const (
	serverPort = 5000
)

// Config is the type used for storing configurations
type Config struct {
	DB         DBConfig
	ServerPort int
	Stage      string
}

// DBConfig stoeres the config for the database
type DBConfig struct {
	User     string
	Name     string
	SSLMode  string
	Host     string
	Password string
}

// Load returns an application config from the file given the current env
func Load(stage string) (*Config, error) {
	var file string
	if !strings.Contains(stage, "/") {
		file = fmt.Sprintf("./config/%s.yml", stage)
	} else {
		file = stage
	}

	c := Config{
		ServerPort: serverPort,
		Stage:      stage,
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
