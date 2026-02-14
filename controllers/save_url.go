package controllers

import (
	"errors"
	"os"
	"os/user"
)

var ErrInvalidPayload = errors.New("invalid payload")

func SaveUrl(payload any) (any, error) {
	url, ok := payload.(string)
	if !ok {
		return nil, ErrInvalidPayload
	}

	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	path := usr.HomeDir + "/url.html"
	content := `<a href="` + url + `">download</a>`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return nil, err
	}

	return map[string]string{"saved": path}, nil
}
