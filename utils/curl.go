package utils

import (
	"log"
	"os/exec"
	"strings"
)

func Curl(bash string, command string) string {
	var _command string
	if len(cfg.Proxy) > 0 {
		proxy := `/usr/bin/curl --proxy "$` + cfg.Proxy + `" `
		_command = strings.Replace(command, "curl", proxy, 1)
	} else {
		_command = strings.Replace(command, "curl", "/usr/bin/curl", 1)
	}

	cmd := exec.Command(bash, "-c", _command)
	out, err := cmd.CombinedOutput()
	// log.Println(string(out))
	if err != nil {
		log.Println(err)
	}
	return string(out)
}
