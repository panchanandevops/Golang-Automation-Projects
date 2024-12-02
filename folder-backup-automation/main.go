package main

import (
	"log"

	"backup-automation/backup"
	"backup-automation/config"
)

func main() {
	// Initialize the backup scheduler
	backupTime := "22:58" // Time to run the backup in HH:MM format
	sourceDir := "/home/panchanan/Pictures/IMG"
	destinationDir := "/home/panchanan/Pictures/Backups"

	err := config.ScheduleBackup(backupTime, func() {
		backup.PerformBackup(sourceDir, destinationDir)
	})
	if err != nil {
		log.Fatalf("Failed to schedule backup: %v", err)
	}

	log.Println("Backup scheduler started...")
	select {} // Keep the main goroutine running
}
