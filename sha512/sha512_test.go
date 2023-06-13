//
// Package sha512
// @Author: feymanlee@gmail.com
// @Description:
// @File:  sha512_test.go
// @Date: 2023/4/23 13:56
//

package sha512

import (
	"crypto/sha512"
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
		{"String", "Hello, world!", "c1527cd893c124773d811911970c8fe6e857d6df5dc9226bd8a160614c0cd963a4ddea2b94bb7d36021ef9d865d5cea294a82dd49a0bb269f51f6e7a57f79421"},
		{"ByteSlice", []byte("Hello, world!"), "c1527cd893c124773d811911970c8fe6e857d6df5dc9226bd8a160614c0cd963a4ddea2b94bb7d36021ef9d865d5cea294a82dd49a0bb269f51f6e7a57f79421"},
		{"Integer", 42, "39ca7ce9ecc69f696bf7d20bb23dd1521b641f806cc7a6b724aaa6cdbffb3a023ff98ae73225156b2c6c9ceddbfc16f5453e8fa49fc10e5d96a3885546a46ef4"},
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

	expected := sha512.Sum512([]byte(content))
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

	expected := sha512.Sum512([]byte(content))
	hexExpected := hex.EncodeToString(expected[:])

	result := MustEncryptFile(tmpfile.Name())
	assert.Equal(t, hexExpected, result)
}
