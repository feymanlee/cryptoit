//
// Package sm3
// @Author: feymanlee@gmail.com
// @Description:
// @File:  sm3_test.go
// @Date: 2023/4/23 13:56
//

package sm3

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjfoc/gmsm/sm3"
)

func TestEncrypt(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"String", "Hello, world!", "e3bca101b496880c3653dad85861d0e784b00a8c18f7574472d156060e9096bf"},
		{"ByteSlice", []byte("Hello, world!"), "e3bca101b496880c3653dad85861d0e784b00a8c18f7574472d156060e9096bf"},
		{"Integer", 42, "3ecd3f08f87cb716ba77980e6aee9ce8b9ce9a61f50d9acd21caa79da8a4bc68"},
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

	h := sm3.New()
	h.Write([]byte(content))
	expected := h.Sum(nil)
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

	h := sm3.New()
	h.Write([]byte(content))
	expected := h.Sum(nil)
	hexExpected := hex.EncodeToString(expected[:])
	result := MustEncryptFile(tmpfile.Name())
	assert.Equal(t, hexExpected, result)
}
