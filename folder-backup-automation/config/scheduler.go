package config

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

// ScheduleBackup schedules a backup task at a specific time.
func ScheduleBackup(backupTime string, task func()) error {
	// Parse the scheduled time
	t, err := time.Parse("15:04", backupTime)
	if err != nil {
		return fmt.Errorf("invalid backup time format: %v", err)
	}

	// Create a scheduler
	scheduler := gocron.NewScheduler(time.Local)

	// Schedule the backup task
	scheduler.Every(1).Day().At(t.Format("15:04")).Do(task)

	// Start the scheduler
	scheduler.StartAsync()
	return nil
}
