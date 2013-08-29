package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) (lines []string, err error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	return
}

func WriteLinesToFile(lines []string, path string) (err error) {
	file, err := os.Create(path)

	if err != nil {
		return nil
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	err = writer.Flush()

	return
}
