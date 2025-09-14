package operations

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"log/slog"
)

var key []byte

func init() {
	// Decode hex string to raw 32-byte key
	k, err := hex.DecodeString("54b431774bce0e0b85989f0ca680e8e9f1ac9f9e6ff6abc795ded8f17e966f00")
	if err != nil {
		log.Fatal("Invalid hex key")
	}
	key = k
}

func encryptAES(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		slog.Error("Encryption key is invalid")
		log.Fatal("Program exiting...")
	}
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	padtext := append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)

	ciphertext := make([]byte, aes.BlockSize+len(padtext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], padtext)
	return ciphertext, nil
}

func decryptAES(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		slog.Error("Invalid key during decryption")
		log.Fatal("Exit 1")
	}
	if len(data) < aes.BlockSize {
		slog.Error("Ciphertext too small")
		log.Fatal("Exit 1")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)

	// unpad
	padding := int(data[len(data)-1])
	return data[:len(data)-padding], nil
}
