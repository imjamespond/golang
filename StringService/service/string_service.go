package service

// StringService provides operations on strings.
import (
	// "context"
	"errors"
	"strings"
)

type IStringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type StringService struct{}

func (StringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (StringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty string")
