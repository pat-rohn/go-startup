package startup

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetLogPath(logfilepath string) error {
	path := filepath.Dir(logfilepath)
	absolutePath, err := filepath.Abs(logfilepath)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v\n", err)
	}
	fmt.Printf("Logfile path is: %s\n", absolutePath)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("Folder does not exist %s: %v\n", path, err)
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatalf("Could not create folder: %v\n", err)
		}
	}
	_, err = os.Stat(logfilepath)
	if os.IsNotExist(err) {
		fmt.Printf("File does not exist %s: %v\n", path, err)
		f, err := os.OpenFile(logfilepath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			log.Fatalf("Could not open log-file: %v\n", err)
		}
		f.Close()
	}
	log.SetOutput(&lumberjack.Logger{
		Filename:   logfilepath,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     100, //days
		Compress:   false,
	})
	fmt.Printf("Log-path set to %s\n", logfilepath)
	return nil
}

func SetLogLevel(logLevel string) {
	switch logLevel {
	case "--trace", "t":
		log.SetLevel(log.TraceLevel)
	case "--info", "i":
		log.SetLevel(log.InfoLevel)
	case "--warn", "w":
		log.SetLevel(log.WarnLevel)
	case "--error", "e":
		log.SetLevel(log.ErrorLevel)
	default:
		fmt.Printf("Invalid log level '%s'\n", logLevel)
	}
	fmt.Printf("LogLevel is set to %s\n", logLevel)
}
