package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Define the username and group to which you want to grant authorization
	username := "your_username"
	groupname := "desired_group_name"

	// Check if the script is executed with root privileges (required to add users to groups)
	if os.Geteuid() != 0 {
		fmt.Println("Please run this script as root or with sudo.")
		os.Exit(1)
	}

	// Check if the group exists; if not, create it
	if !groupExists(groupname) {
		if err := createGroup(groupname); err != nil {
			fmt.Printf("Error creating group '%s': %v\n", groupname, err)
			os.Exit(1)
		}
		fmt.Printf("Group '%s' has been created.\n", groupname)
	}

	// Add the user to the specified group
	if err := addUserToGroup(username, groupname); err != nil {
		fmt.Printf("Error adding user '%s' to group '%s': %v\n", username, groupname, err)
		os.Exit(1)
	}

	fmt.Printf("User '%s' has been granted authorization in group '%s'.\n", username, groupname)
}

func groupExists(groupname string) bool {
	cmd := exec.Command("getent", "group", groupname)
	return cmd.Run() == nil
}

func createGroup(groupname string) error {
	cmd := exec.Command("groupadd", groupname)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func addUserToGroup(username, groupname string) error {
	cmd := exec.Command("usermod", "-aG", groupname, username)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
