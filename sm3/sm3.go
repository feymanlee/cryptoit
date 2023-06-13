//
// Package sm3
// @Author: feymanlee@gmail.com
// @Description:
// @File:  sm3
// @Date: 2023/4/23 13:55
//

package sm3

import (
	"encoding/hex"
	"io"
	"os"

	"github.com/feymanlee/cryptoit"
	"github.com/tjfoc/gmsm/sm3"
)

// Encrypt encrypts any type of variable using SHA1 algorithms.
func Encrypt(v interface{}) string {
	h := sm3.New()
	h.Write(cryptoit.AnyToByte(v))
	return hex.EncodeToString(h.Sum(nil))
}

// EncryptFile encrypts file content of `path` using SHA1 algorithms.
func EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sm3.New()
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
