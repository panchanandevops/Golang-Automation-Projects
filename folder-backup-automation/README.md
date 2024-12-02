# Folder Backup Automation

This is a simple folder backup automation tool written in Go. The application schedules a daily backup of a specified source folder to a destination folder at a specified time. It copies the contents of the source folder into a new folder, named with the current date, in the destination folder.

## Project Structure

```
folder-backup-automation
├── backup
│   ├── backup.go
│   └── copier.go
├── config
│   └── scheduler.go
├── go.mod
├── go.sum
└── main.go
```

### **Main Files:**
- **`main.go`**: The entry point of the application where the backup scheduler is initialized and executed.
- **`backup/backup.go`**: Contains the `PerformBackup` function that handles the backup operation.
- **`backup/copier.go`**: Provides functions to copy files and folders from the source directory to the destination directory.
- **`config/scheduler.go`**: Handles the scheduling of the backup task at a specific time.

## Installation

1. **Install dependencies:**
   The project uses Go modules to manage dependencies. Run the following to download the necessary dependencies:
   ```bash
   go mod tidy
   ```

## How to Run

1. **Run the application:**
   Start the backup scheduler by running:
   ```bash
   go run main.go
   ```

   The application will continuously run in the background and will execute the backup task at the scheduled time (22:58 in the example below).

2. **Configure your directories and time:**
   In the `main.go` file, modify the `sourceDir`, `destinationDir`, and `backupTime` to suit your requirements:
   ```go
   backupTime := "22:58" // Time to run the backup in HH:MM format
   sourceDir := "/home/user/source"  // Path to the source directory
   destinationDir := "/home/user/backup"  // Path to the destination directory
   ```

3. **Backup operation:**
   The tool will create a backup folder in the destination directory with the current date as the folder name (e.g., `/home/user/backup/2024-12-02`).

## How It Works

### **1. Scheduling the Backup**
The `ScheduleBackup` function, defined in `config/scheduler.go`, uses the **gocron** package to schedule a backup task at a specified time. The backup is triggered daily at the provided time (e.g., 22:58).

### **2. Performing the Backup**
The `PerformBackup` function in `backup/backup.go` performs the backup task by:
- Creating a new folder in the destination directory named with today's date.
- Calling the `CopyFolder` function to recursively copy the contents of the source directory into the new backup folder.

### **3. Copying Files and Folders**
The `CopyFolder` function in `backup/copier.go` copies the entire folder structure from the source to the destination. If a directory or file already exists, it is overwritten. The files are copied using the `CopyFile` function.

## Dependencies

- **Go**: The Go programming language is required to build and run the project.
- **gocron**: A simple cron job library for Go, used to schedule the backup tasks.
  
You can install the dependencies by running:
```bash
go get github.com/go-co-op/gocron
```

