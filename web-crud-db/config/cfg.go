package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadProperties(filename string) {

	f, err := os.Open(".env")
	if err != nil {
		log.Fatalf("%s not found", filename)
	}

	reader := bufio.NewScanner(f)

	for reader.Scan() {
		params := strings.Split(reader.Text(), "=")
		os.Setenv(params[0], params[1])
	}

}
