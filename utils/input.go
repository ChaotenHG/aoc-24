package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLine(path string) []string {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error() + `: ` + path)
		return []string{}
	}
	defer inFile.Close()

	lines := []string{}

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
