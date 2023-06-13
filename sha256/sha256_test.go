//
// Package sha256
// @Author: feymanlee@gmail.com
// @Description:
// @File:  sha256_test.go
// @Date: 2023/4/23 13:56
//

package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"String", "Hello, world!", "315f5bdb76d078c43b8ac0064e4a0164612b1fce77c869345bfc94c75894edd3"},
		{"ByteSlice", []byte("Hello, world!"), "315f5bdb76d078c43b8ac0064e4a0164612b1fce77c869345bfc94c75894edd3"},
		{"Integer", 42, "73475cb40a568e8da8a045ced110137e159f890ac4da883b6b17dc651b3a8049"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Encrypt(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestEncryptFile(t *testing.T) {
	content := "Hello, world!"
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	expected := sha256.Sum256([]byte(content))
	hexExpected := hex.EncodeToString(expected[:])

	result, err := EncryptFile(tmpfile.Name())
	assert.NoError(t, err)
	assert.Equal(t, hexExpected, result)
}

func TestMustEncryptFile(t *testing.T) {
	content := "Hello, world!"
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MustEncryptFile() panics: %v", r)
		}
	}()

	expected := sha256.Sum256([]byte(content))
	hexExpected := hex.EncodeToString(expected[:])

	result := MustEncryptFile(tmpfile.Name())
	assert.Equal(t, hexExpected, result)
}
