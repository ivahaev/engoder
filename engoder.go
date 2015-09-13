package engoder

import (
	"bytes"
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/scrypt"
)

// Encrypt takes input string and options []byte salt and return encrypted hash, salt and error
func Encrypt(input string, _salt ...[]byte) (hash []byte, salt []byte, err error) {
	if len(_salt) > 0 {
		salt = _salt[0]
	} else {
		salt = make([]byte, 16)
		_, err = rand.Read(salt[:16])
		if err != nil {
			return nil, nil, errors.New("Can't generate salt")
		}
	}
	hash, err = scrypt.Key([]byte(input), salt, 16384, 8, 1, 32)
	if err != nil {
		return nil, nil, errors.New("Can't encrypt input")
	}
	return hash, salt, nil
}

// Check will encrypt input string with salt provided and compare with _hash. If equals, will return true.
func Check(input string, _hash, salt []byte) bool {
	hash, _, err := Encrypt(input, salt)
	if err != nil {
		return false
	}
	return bytes.Equal(hash, _hash)
}
