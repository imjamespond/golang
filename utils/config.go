package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Bash     string
	Proxy    string
	Curl     string
	Interval int
}

const cfgYaml = ".config.yaml"

var cfg = Config{}

func GetCfg() *Config {
	return &cfg
}

func isExist(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return !info.IsDir()
}

func ReadCfg() {
	if !isExist(cfgYaml) {
		cfg, err := os.Create(cfgYaml)
		if err != nil {
			log.Fatal(err)
		}
		cfg.Close()
	}

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
