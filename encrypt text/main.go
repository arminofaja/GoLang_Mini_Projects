package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func generateAESKey() ([]byte, error) {
	key := make([]byte, 32) // 256 bits
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func encryptAES(key []byte, plaintext []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

func main() {
	plaintext := []byte("Hello, World!") // Write the plaintext

	// Generate a random AES key (you should handle key management securely)
	key, err := generateAESKey()
	if err != nil {
		fmt.Println("Key generation error:", err)
		return
	}

	ciphertext, nonce, err := encryptAES(key, plaintext)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}

	fmt.Println("Plaintext:", string(plaintext))
	fmt.Printf("Ciphertext (hex): %x\n", ciphertext)
	fmt.Printf("Nonce (hex): %x\n", nonce)
}
