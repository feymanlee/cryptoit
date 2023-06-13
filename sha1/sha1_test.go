//
// Package sha1
// @Author: feymanlee@gmail.com
// @Description:
// @File:  sha1_test.go
// @Date: 2023/4/23 13:51
//

package sha1

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected string
	}{
		{"hello", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"},
		{[]byte("hello"), "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"},
		{123, "40bd001563085fc35165329ea1ff5c5ecbdbbeef"},
	}

	for _, tc := range testCases {
		result := Encrypt(tc.input)
		assert.Equal(t, tc.expected, result)
	}
}

func TestEncryptFile(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "test_file")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString("hello")
	assert.NoError(t, err)
	err = tempFile.Close()
	assert.NoError(t, err)

	result, err := EncryptFile(tempFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", result)

	_, err = EncryptFile("non_existent_file")
	assert.Error(t, err)
}

func TestMustEncryptFile(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "test_file")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString("hello")
	assert.NoError(t, err)
	err = tempFile.Close()
	assert.NoError(t, err)

	assert.NotPanics(t, func() {
		result := MustEncryptFile(tempFile.Name())
		assert.Equal(t, "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", result)
	})

	assert.Panics(t, func() {
		MustEncryptFile("non_existent_file")
	})
}
