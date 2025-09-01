package dataIO

import (
	"bufio"
	"io"
	"os"
)

func Read(inputFile string) ([]string, error) {
	if inputFile == "" {
		return read(os.Stdin)
	}
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	return read(f)
}

func read(reader io.Reader) ([]string, error) {
	var data []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, scanner.Err()
}

// Write(outputFile string, data []string)

// write(writer io.Writer, data []string)
