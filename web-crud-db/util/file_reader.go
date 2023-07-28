package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) map[string]string {

	readFile, err := os.Open(filename)

	if err != nil {
		log.Fatal("ReadFile= " + err.Error())
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	defer readFile.Close()

	props := make(map[string]string)

	for _, line := range fileLines {
		s := strings.Split(line, "=")
		props[s[0]] = s[1]
	}

	return props
}
