package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateCaptcha() string {
	rand.Seed(time.Now().UnixNano())
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	captchaLength := 6 // You can adjust the CAPTCHA length as needed

	var captcha string
	for i := 0; i < captchaLength; i++ {
		randomIndex := rand.Intn(len(characters))
		captcha += string(characters[randomIndex])
	}

	return captcha
}

func main() {
	captcha := generateCaptcha()

	fmt.Println("Welcome to the Human Verification Program!")
	fmt.Println("Please enter the following CAPTCHA:")
	fmt.Println(captcha)

	var userInput string
	fmt.Print("Enter CAPTCHA: ")
	fmt.Scanln(&userInput)

	if userInput == captcha {
		fmt.Println("Verification successful. You are a human!")
	} else {
		fmt.Println("Verification failed. You might be a bot.")
	}

	// Wait for user input before exiting
	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}
