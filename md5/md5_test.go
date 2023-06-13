package md5

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	data := []byte("Hello, world!")
	dataStr := "Hello, world!"
	expected := fmt.Sprintf("%x", md5.Sum(data))

	// Encrypt
	encrypted, err := Encrypt(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, encrypted)

	// MustEncrypt
	assert.NotPanics(t, func() {
		result := MustEncrypt(data)
		assert.Equal(t, expected, result)
	})

	// EncryptBytes
	encryptedBytes, err := EncryptBytes(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, encryptedBytes)

	// MustEncryptBytes
	assert.NotPanics(t, func() {
		result := MustEncryptBytes(data)
		assert.Equal(t, expected, result)
	})

	// EncryptString
	encryptedStr, err := EncryptString(dataStr)
	assert.NoError(t, err)
	assert.Equal(t, expected, encryptedStr)

	// MustEncryptString
	assert.NotPanics(t, func() {
		result := MustEncryptString(dataStr)
		assert.Equal(t, expected, result)
	})

	// EncryptFile and MustEncryptFile
	tmpfile, err := ioutil.TempFile("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write(data)
	assert.NoError(t, err)
	err = tmpfile.Close()
	assert.NoError(t, err)

	encryptedFile, err := EncryptFile(tmpfile.Name())
	assert.NoError(t, err)
	assert.Equal(t, expected, encryptedFile)

	assert.NotPanics(t, func() {
		result := MustEncryptFile(tmpfile.Name())
		assert.Equal(t, expected, result)
	})
}
