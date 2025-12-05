package io

import (
	"bufio"
	"os"
)

func ExtractCommandsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var commands []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := scanner.Text()
		commands = append(commands, command)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}
