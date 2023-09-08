package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	// Check if the script is executed with root privileges
	if os.Geteuid() != 0 {
		fmt.Println("Please run this script as root or with sudo.")
		os.Exit(1)
	}

	// Prompt for the username
	fmt.Print("Enter the username: ")
	var username string
	fmt.Scanln(&username)

	// Check if the username already exists
	if _, err := user.Lookup(username); err == nil {
		fmt.Printf("User '%s' already exists.\n", username)
		os.Exit(1)
	}

	// Prompt for the user's full name
	fmt.Print("Enter the full name of the user: ")
	var fullName string
	fmt.Scanln(&fullName)

	// Prompt for and set the user's password
	fmt.Print("Enter the password: ")
	var password string
	fmt.Scanln(&password)

	// Create the user with the provided username and full name
	cmd := exec.Command("useradd", "-c", fullName, "-m", username)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to create user '%s': %v\n", username, err)
		os.Exit(1)
	}

	// Set the user's password
	cmd = exec.Command("passwd", username)
	cmd.Stdin = strings.NewReader(password + "\n" + password + "\n") // Set the same password twice
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to set a password for user '%s': %v\n", username, err)
		os.Exit(1)
	}

	fmt.Printf("User '%s' has been created successfully.\n", username)

}

//go build add_user_no_terminal.go   and always use sudo    sudo ./add_user_no_terminal
