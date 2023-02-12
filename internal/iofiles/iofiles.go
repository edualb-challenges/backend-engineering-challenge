package iofiles

import (
	"bufio"
	"fmt"
	"os"
)

// GetFirstLine returns the first line of a file
func GetFirstLine(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	firstLine := scanner.Bytes()

	if string(firstLine) == "" {
		return nil, fmt.Errorf("empty file")
	}

	return firstLine, nil
}

// GetLastLine returns the last line of a file
func GetLastLine(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := stat.Size()

	buf := make([]byte, fileSize)
	file.Read(buf)

	err = file.Close()
	if err != nil {
		return nil, err
	}

	startByte := 0
	for i := len(buf) - 1; i >= 0; i-- {
		if buf[i] == '\n' {
			break
		}
		startByte = i
	}

	lastLine := buf[startByte:]

	if string(lastLine) == "" {
		return nil, fmt.Errorf("empty file")
	}

	return lastLine, nil
}
