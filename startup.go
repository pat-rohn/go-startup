package startup

import (
	"fmt"
	"os"

	"picloud.ch/schusti/pkg/config"

	log "github.com/sirupsen/logrus"
)

const (
	logPkg string = "startup"
)

func SetLogLevel(logLevel string, path string) bool {
	switch logLevel {
	case "--trace", "-t":
		log.SetLevel(log.TraceLevel)
		break
	case "--info", "-i":
		log.SetLevel(log.InfoLevel)
		break
	case "--warn", "-w":
		log.SetLevel(log.WarnLevel)
		break
	case "--error", "-e":
		log.SetLevel(log.ErrorLevel)
		break
	case "--logfile", "-f":
		fmt.Println("Check if file exist")
		fullPath := config.LogPath + path
		err := os.Chmod(config.LogPath, 0777)
		if err != nil {
			fmt.Printf("Could not change permission rights %s: %v\n", path, err)
			os.Exit(0)
		}
		if _, err := os.Stat(fullPath); err == nil {
			fmt.Printf("File exist: %v\n", err)
			//backupFile := config.LogPath + time.Now().Format(driverinterface.TimestampFormatFilename+path)
			//fmt.Printf("Backup to file %s\n", backupFile)

			err := os.Rename(fullPath, fullPath+"2")
			if err != nil {
				fmt.Printf("Could not rename log-file %s: %v\n", path, err)
				os.Exit(0)
			}
		}
		f, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Printf("Could not open log-file: %v\n", err)
			os.Exit(0)
		}
		log.SetOutput(f)
		fmt.Printf("Logging set to %s\n", fullPath)
		return false
	default:
		return false
	}
	fmt.Printf("LogLevel is set to %s\n", logLevel)
	return true
}
