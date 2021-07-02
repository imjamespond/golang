package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ParseConfig(file string) map[string]interface{} {
	cfgData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(string(cfgData))

	cfg := make(map[string]interface{})
	if err := json.Unmarshal(cfgData, &cfg); err != nil {
		log.Fatal(err)
	}
	// log.Println(cfg)
	return cfg
}
