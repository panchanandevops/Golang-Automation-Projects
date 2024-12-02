package backup

import (
	"log"
	"path/filepath"
	"time"
)

// PerformBackup performs the backup operation from source to destination.
func PerformBackup(sourceDir, destinationDir string) {
	// Get today's date for the backup folder
	today := time.Now().Format("2006-01-02")
	backupDir := filepath.Join(destinationDir, today)

	err := CopyFolder(sourceDir, backupDir)
	if err != nil {
		log.Printf("Error during backup: %v\n", err)
	} else {
		log.Printf("Backup completed successfully: %s\n", backupDir)
	}
}
