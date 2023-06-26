package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

var aesKey = []byte("too young too simple too naive!!")

func EncryptWithAES(plainText string) (string, error) {
	// 创建新的Cipher块
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	plainTextBytes := []byte(plainText)
	cipherText := make([]byte, aes.BlockSize+len(plainTextBytes))

	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainTextBytes)

	return hex.EncodeToString(cipherText), nil
}

func DecryptWithAES(cipherText string) (string, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	decodedCipherText, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	if len(decodedCipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := decodedCipherText[:aes.BlockSize]
	decodedCipherText = decodedCipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decodedCipherText, decodedCipherText)

	return string(decodedCipherText), nil
}
