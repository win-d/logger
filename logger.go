// Package logger provides simple functions for writing to log files
package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// dir is the directory where we create the log files
var dir = "/var/log/project"

// SetDir sets (or creates) a directory for log files
func SetDir(directory string) error {
	directory = filepath.FromSlash(directory)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err = createDir(directory); err != nil {
			return err
		}
	}

	dir = directory
	return nil
}

// Write writes a text to the log file
func Write(text string) {
	mutex := &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	file, err := openFile()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.SetOutput(file)
	log.Println(text)

	err = closeFile(file)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// createDir creates a directory for log files
func createDir(directory string) error {
	err := os.Mkdir(directory, 0755)
	if err != nil {
		return err
	}

	return nil
}

// openFile opens (or creates) a log file. If successful,
// method on the returned File can be used for I/O.
func openFile() (*os.File, error) {
	var (
		sep   = string(os.PathSeparator)
		today = time.Now()
	)

	directory := dir + sep + fmt.Sprintf("%02d", today.Month())
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err = createDir(directory); err != nil {
			return nil, err
		}
	}

	name := directory + sep + today.Format("02012006") + ".log"
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// closeFile closes the File, rendering it unusable for I/O
func closeFile(file *os.File) error {
	if err := file.Close(); err != nil {
		return err
	}

	return nil
}
