package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	inputfile := "standard.txt"

	file, err := os.ReadFile(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	word := string(file)
	char := 'A'

	splitfile := strings.Split(word, "\n")
	start := (int(char) - 32) * 8
	var result rune
	for val := range char {
		for i := start; i < 8; i++ {
			result += val
		}
	}

	os.Stdout.WriteString(string(file))
}
