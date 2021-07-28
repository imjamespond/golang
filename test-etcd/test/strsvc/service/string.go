package service

import (
	"strings"
	"test-etcd/test/strsvc/util"
)

// StringService provides operations on strings.
type IStringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type StringService struct{}

func (StringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", util.ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (StringService) Count(s string) int {
	return len(s)
}
