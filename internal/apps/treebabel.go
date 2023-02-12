package apps

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/edualb-challenge/treebabel/internal/models"
)

// TreeBabel is an application responsible to process the average delivery time
type TreeBabel struct {
	TranslationDeliveredFile string
	WindowSize               uint64
}

// Run executes the treebabel application
func (app TreeBabel) Run() error {
	firstTime, err := getFirstTime(app.TranslationDeliveredFile)
	if err != nil {
		return err
	}
	lastTime, err := getLastTime(app.TranslationDeliveredFile)
	if err != nil {
		return err
	}
	fmt.Printf("First Hour: %s\n", firstTime.String())
	fmt.Printf("Last Hour: %s\n", lastTime.String())
	return nil
}

// getFirstTime is responsible to get the first time located at the first register of the input file
func getFirstTime(filePath string) (time.Time, error) {
	var firstTime time.Time
	var firstTD models.TranslationDelivered

	file, err := os.Open(filePath)
	if err != nil {
		return firstTime, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	err = json.Unmarshal(scanner.Bytes(), &firstTD)
	if err != nil {
		return firstTime, err
	}
	firstTime = firstTD.Timestamp.Time

	err = scanner.Err()
	if err != nil {
		return firstTime, err
	}

	err = file.Close()
	if err != nil {
		return firstTime, err
	}

	firstTime = time.Date(firstTime.Year(), firstTime.Month(), firstTime.Day(), firstTime.Hour(), firstTime.Minute(), 0, 0, firstTime.Location())
	return firstTime, nil
}

// getLastTime is responsible to get the last time located at the last register of the input file
func getLastTime(filePath string) (time.Time, error) {
	var lastTime time.Time
	var lastTD models.TranslationDelivered

	file, err := os.Open(filePath)
	if err != nil {
		return lastTime, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return lastTime, err
	}
	fileSize := stat.Size()

	buf := make([]byte, fileSize)
	file.Read(buf)

	startByte := 0
	for i := len(buf) - 1; i >= 0; i-- {
		if buf[i] == '\n' {
			break
		}
		startByte = i
	}

	err = json.Unmarshal(buf[startByte:], &lastTD)
	if err != nil {
		return lastTime, err
	}
	lastTime = lastTD.Timestamp.Time

	minute := lastTime.Minute()
	if lastTime.Second() > 0 {
		minute = lastTime.Minute() + 1
	}

	lastTime = time.Date(lastTime.Year(), lastTime.Month(), lastTime.Day(), lastTime.Hour(), minute, 0, 0, lastTime.Location())

	return lastTime, nil
}
