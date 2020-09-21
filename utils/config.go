package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	configName = "application.yml"
)

type Config struct {
	Rpc
	Registry
	Mysql
	Redis
	Gateway
	Im
	Version string
}

type Gateway struct {
	List []string `yaml:",flow"`
	Port int
}

type Rpc struct {
	Port     int
	Server   string
	Metadata string
}

type Registry struct {
	Adders []string `yaml:",flow"`
}

type Mysql struct {
	User     string
	Password string
	Db       string
}

type Im struct {
	Index string
	Port  int
}

type Redis struct {
}

func LoadConfig() (*Config, error) {
	conf := new(Config)
	yamlFile, err := ioutil.ReadFile(configName)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(yamlFile, conf); err != nil {
		panic(err)
	}
	return conf, nil
}

func ScanConfig(c *Config) error {
	yamlFile, err := ioutil.ReadFile(configName)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		panic(err)
	}
	return nil
}
