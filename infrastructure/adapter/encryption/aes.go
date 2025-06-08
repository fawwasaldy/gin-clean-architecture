package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"kpl-base/domain/port"
	"os"
)

type aesAdapter struct{}

func NewAesAdapter() port.EncryptionPort {
	return &aesAdapter{}
}

func (a aesAdapter) Encrypt(plainText string) (cipherText string, err error) {
	key, err := hex.DecodeString(os.Getenv("AES_KEY"))
	if err != nil {
		return "", err
	}
	plainTextByte := []byte(plainText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherTextByte := aesGCM.Seal(nonce, nonce, plainTextByte, nil)
	return fmt.Sprintf("%x", cipherTextByte), nil
}

func (a aesAdapter) Decrypt(cipherText string) (plainText string, err error) {
	defer func() {
		if r := recover(); r != nil {
			plainText = ""
			err = errors.New("decryption failed")
		}
	}()

	key, err := hex.DecodeString(os.Getenv("AES_KEY"))
	if err != nil {
		return "", err
	}

	cipherTextByte, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	nonce, cipherTextByte := cipherTextByte[:nonceSize], cipherTextByte[nonceSize:]

	plainTextByte, err := aesGCM.Open(nil, nonce, cipherTextByte, nil)
	if err != nil {
		return "", err
	}

	return string(plainTextByte), nil
}
