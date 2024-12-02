package backup

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyFolder copies all files and directories from the source to the destination.
func CopyFolder(source, dest string) error {
	// Create the destination folder if it doesn't exist
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		err = os.MkdirAll(dest, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create destination directory: %v", err)
		}
	}

	// Walk through the source directory and copy files
	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %s: %v", path, err)
		}

		// Create the destination path
		relativePath, err := filepath.Rel(source, path)
		if err != nil {
			return fmt.Errorf("error calculating relative path: %v", err)
		}
		destPath := filepath.Join(dest, relativePath)

		if info.IsDir() {
			// Create the directory
			return os.MkdirAll(destPath, os.ModePerm)
		} else {
			// Copy the file
			return CopyFile(path, destPath)
		}
	})

	return err
}

// CopyFile copies a single file from the source to the destination.
func CopyFile(source, dest string) error {
	srcFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file contents: %v", err)
	}

	return nil
}
