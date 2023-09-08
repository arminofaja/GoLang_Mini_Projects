package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

const (
	passwordLength  = 16 // Adjust the length of the generated password as needed
	passwordCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>?`~"
	passwordFile    = "Data/password.txt" // File to save the generated password in a 'data' directory relative to the program
)

func generatePassword(length int, charset string) (string, error) {
	var password strings.Builder

	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		randomChar := charset[randomIndex.Int64()]
		password.WriteByte(randomChar)
	}

	return password.String(), nil
}

func savePasswordToFile(password string, filename string) error {
	// Ensure the directory exists
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// Open the file in append mode and create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(password + "\n")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	password, err := generatePassword(passwordLength, passwordCharset)
	if err != nil {
		fmt.Println("Password generation error:", err)
		return
	}

	fmt.Println("Generated Password:", password)

	err = savePasswordToFile(password, passwordFile)
	if err != nil {
		fmt.Println("Error saving password to file:", err)
		return
	}

	fmt.Printf("Password appended to %s\n", passwordFile)
}
