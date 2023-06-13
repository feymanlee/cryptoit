//
// Package sha512
// @Author: feymanlee@gmail.com
// @Description:
// @File:  sha512
// @Date: 2023/4/23 13:55
//

package sha512

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"

	"github.com/feymanlee/cryptoit"
)

// Encrypt encrypts any type of variable using SHA1 algorithms.
func Encrypt(v interface{}) string {
	r := sha512.Sum512(cryptoit.AnyToByte(v))
	return hex.EncodeToString(r[:])
}

// EncryptFile encrypts file content of `path` using SHA1 algorithms.
func EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha512.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// MustEncryptFile encrypts file content of `path` using SHA1 algorithms.
// It panics if any error occurs.
func MustEncryptFile(path string) string {
	result, err := EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
