package utils

import (
	"fmt"
	"log"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type YamlTest struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func TestYaml(_t *testing.T) {

	t := YamlTest{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))

}

//判断文件或文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		log.Println(err.Error())

		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func TestYaml1(t *testing.T) {
	cfgYaml := ".config.yaml"
	// if !IsExist(cfgYaml) {
	// 	cfg, err := os.Create(cfgYaml)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(cfg)
	// 	cfg.Close()
	// 	// cfgData, err := ioutil.ReadAll(cfg)
	// }

	cfgData, err := os.ReadFile(cfgYaml)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(cfgData))

	// yaml

	cfg := Config{}

	// yaml read
	err = yaml.Unmarshal([]byte(cfgData), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(cfg)

	cfg.Bash = "D:/app/msys 64/usr/bin/bash.exe"

	// yaml write
	yamlData, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(string(yamlData))

	// write cfg into file
	/*
		WriteFile writes data to the named file, creating it if necessary.
		If the file does not exist, WriteFile creates it with permissions perm
	*/
	os.WriteFile(cfgYaml, yamlData, 0611)
}
