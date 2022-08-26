package simple

import (
	"crypto/rand"
	"encoding/base64"
)

func Password() (password string, err error) {
	var buf [16]byte

	_, err = rand.Read(buf[:])
	password = base64.RawStdEncoding.EncodeToString(buf[:])

	return
}

