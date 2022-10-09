package main

import (
	"crypto/rand"
	"encoding/base64"
)

type simple struct{}

func (s simple) Password() (password string, err error) {
	var buf [16]byte

	_, err = rand.Read(buf[:])
	password = base64.RawStdEncoding.EncodeToString(buf[:])

	return
}

