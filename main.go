package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

var fileMap = make(map[string]bool)

func scanDirectory(source string, target string, rename bool) {
	files, err := os.ReadDir(source)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileMap[file.Name()] = true
		log.Printf("New file found: %s\n", file.Name())
		id := uuid.New()

		log.Printf("New job inserted with id: %s\n", id)

		f, err := os.Open(fmt.Sprintf("%s/%s", source, file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		fileType := splitFileName(file.Name())

		newFileName := fmt.Sprintf("%s/%s", target, file.Name())
		if rename {
			if fileType != nil {
				newFileName = fmt.Sprintf("%s/%s.%s", target, id, *fileType)
			} else {
				newFileName = fmt.Sprintf("%s/%s", target, id)
			}
		}

		log.Printf("Copying file %s to %s\n", file.Name(), newFileName)

		newFile, err := os.Create(newFileName)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.WriteTo(newFile)
		if err != nil {
			log.Fatal(err)
		}

		err = os.Remove(fmt.Sprintf("%s/%s", source, file.Name()))
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("File %s copied to %s\n", file.Name(), newFileName)
	}
}

func splitFileName(fileName string) *string {
	fileNameSplit := strings.Split(fileName, ".")
	if len(fileNameSplit) == 1 {
		return nil
	}

	return &fileNameSplit[len(fileNameSplit)-1]
}

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Usage: program <directory_path> <timer_duration>")
	}

	sourceDir := os.Args[1]
	targetDir := os.Args[2]
	timer := os.Args[3]
	timerInt, err := time.ParseDuration(timer)
	if err != nil {
		log.Fatal(err)
	}

	scanDirectory(sourceDir, targetDir, false)

	ticker := time.NewTicker(timerInt)
	defer ticker.Stop()

	for range ticker.C {
		scanDirectory(sourceDir, targetDir, false)
	}
}
