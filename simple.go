package main

import (
	"crypto/rand"
	"encoding/base64"
)

const defaultSimplePasswordLength = 16

type simple struct {
	length int
}

func (s simple) Password() (password string, err error) {
	length := defaultSimplePasswordLength

	if s.length > 0 {
		length = s.length
	}

	buf := make([]byte, length)

	_, err = rand.Read(buf[:])
	password = base64.RawStdEncoding.EncodeToString(buf[:])[:length]

	return
}
