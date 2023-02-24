package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	return hex.EncodeToString(bs)
}
