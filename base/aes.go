package base

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

// AES CBC
func aesEncryptCBC(origData, key []byte) (crypted []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			crypted = nil
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	blockSize := block.BlockSize()
	origData = zeroPadding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted = make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return
}

// AES CBC Do a Base64 encryption after encryption
func aesEncryptCBCWithBase64(origData, key []byte) (string, error) {
	cbc, err := aesEncryptCBC(origData, key)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cbc), nil
}

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	if padding == 0 {
		return ciphertext
	}

	padtext := bytes.Repeat([]byte{byte(0)}, padding)
	return append(ciphertext, padtext...)
}
