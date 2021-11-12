package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Bash string
}

const cfgYaml = ".config.yaml"

var cfg = Config{}

func init() {
	ReadCfg()
}

func GetCfg() *Config {
	return &cfg
}

func ReadCfg() {
	cfgData, err := os.ReadFile(cfgYaml)
	if err != nil {
		log.Println(err)
	}

	// yaml read
	err = yaml.Unmarshal([]byte(cfgData), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(cfg)
}
func WriteCfg() {
	// yaml write
	yamlData, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(string(yamlData))
	os.WriteFile(cfgYaml, yamlData, 0611)
}
