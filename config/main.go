package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//AppConfig is the main config file structure
type AppConfig struct {
	Server    ServerConfig `yaml:"server"`
	DbConfigs []DbConfig   `yaml:"dbs"`
}

//ServerConfig represents the Server configuration
type ServerConfig struct {
	Port uint
}

//DbConfig is a single DataBase connection config
type DbConfig struct {
	Host   string
	DbName string
	Port   string
	User   string
	Pass   string
}

func loadCfgFromDisk() ([]byte, error) {
	return ioutil.ReadFile("db.yaml")
}

//GetConfig loads the app config
func GetConfig() AppConfig {
	conf := AppConfig{}
	data, err := loadCfgFromDisk()
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, &conf); err != nil {
		panic(err)
	}

	return conf
}
