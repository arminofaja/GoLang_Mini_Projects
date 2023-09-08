package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Define source and destination folders
	sourceFolder := "/path/to/source_folder"
	destinationFolder := "/path/to/destination_folder"

	// Create a timestamp for the index file
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	indexFilename := fmt.Sprintf("log_index_%s.txt", timestamp)

	// Create the index file
	indexFile, err := os.Create(filepath.Join(destinationFolder, indexFilename))
	if err != nil {
		fmt.Printf("Error creating index file: %v\n", err)
		return
	}
	defer indexFile.Close()

	// List files in the source folder
	files, err := ioutil.ReadDir(sourceFolder)
	if err != nil {
		fmt.Printf("Error listing files in source folder: %v\n", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			sourceFilePath := filepath.Join(sourceFolder, file.Name())
			destinationFilePath := filepath.Join(destinationFolder, file.Name())

			// Copy the file to the destination folder
			if err := copyFile(sourceFilePath, destinationFilePath); err != nil {
				fmt.Printf("Error copying %s: %v\n", file.Name(), err)
				continue
			}

			// Write the file name and date to the index file
			indexFile.WriteString(fmt.Sprintf("%s: %s\n", file.Name(), time.Now().Format(time.RFC3339)))
		}
	}

	fmt.Println("Log files copied successfully.")
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
