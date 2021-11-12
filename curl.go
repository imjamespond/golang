package main

import (
	"log"
	"os/exec"
)

func curl(bash string, command string) {
	cmd := exec.Command(bash, "-c", "/usr/bin/"+command)
	out, err := cmd.CombinedOutput()
	log.Println(string(out))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
