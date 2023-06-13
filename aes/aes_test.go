//
// Package aes
// @Author: feymanlee@gmail.com
// @Description:
// @File:  aes_test.go
// @Date: 2023/4/23 15:55
//

package aes

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAES(t *testing.T) {
	key := []byte("0123456789abcdef")
	iv := []byte("abcdef0123456789")
	plainText := []byte("Hello, world!")

	// Encrypt
	encrypted, err := Encrypt(plainText, key, iv)
	assert.NoError(t, err)

	// Decrypt
	decrypted, err := Decrypt(encrypted, key, iv)
	assert.NoError(t, err)
	assert.Equal(t, plainText, decrypted)

	// EncryptCBC
	encryptedCBC, err := EncryptCBC(plainText, key, iv)
	assert.NoError(t, err)
	assert.Equal(t, encrypted, encryptedCBC)

	// DecryptCBC
	decryptedCBC, err := DecryptCBC(encryptedCBC, key, iv)
	assert.NoError(t, err)
	assert.Equal(t, plainText, decryptedCBC)

	// PKCS5Padding
	paddedText := PKCS5Padding(plainText, aes.BlockSize)

	// PKCS5UnPadding
	unpaddedText, err := PKCS5UnPadding(paddedText, aes.BlockSize)
	assert.NoError(t, err)
	assert.Equal(t, plainText, unpaddedText)

	// EncryptCFB
	var padding int
	encryptedCFB, err := EncryptCFB(plainText, key, &padding, iv)
	assert.NoError(t, err)

	// DecryptCFB
	decryptedCFB, err := DecryptCFB(encryptedCFB, key, padding, iv)
	assert.NoError(t, err)
	assert.Equal(t, plainText, decryptedCFB)

	// ZeroPadding
	paddedTextZero, zeroPadding := ZeroPadding(plainText, aes.BlockSize)

	// ZeroUnPadding
	unpaddedTextZero := ZeroUnPadding(paddedTextZero, zeroPadding)
	assert.Equal(t, plainText, unpaddedTextZero)
}
