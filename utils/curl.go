package utils

import (
	"log"
	"os/exec"
	"strings"
)

func Curl(bash string, command string) string {
	var _command string
	curl := "/usr/bin/curl"
	if len(cfg.Curl) > 0 {
		curl = cfg.Curl
	}
	if len(cfg.Proxy) > 0 {
		proxy := curl + " --proxy " + cfg.Proxy
		_command = strings.Replace(command, "curl", proxy, 1)
	} else {
		_command = strings.Replace(command, "curl", curl, 1)
	}
	log.Println(_command)
	cmd := exec.Command(bash, "-c", _command)
	out, err := cmd.CombinedOutput()
	// log.Println(string(out))
	if err != nil {
		log.Println(err)
	}
	return string(out)
}
