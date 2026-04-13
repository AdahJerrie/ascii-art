package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) >= 2 {

	}
	bannerfile := "standard.txt"

	text := os.Args[1]
	if len(os.Args) < 2 {
		fmt.Println("Invalid Arg length")
		return
	}

	file, err := os.ReadFile(bannerfile)
	if err != nil {
		log.Fatal(err)
	}
	content := string(file)

	lines := strings.Split(content, "\n") // split banner files by newlines
	textlines := strings.Split(text, "\\n")
	var result string

	for _, val := range textlines {
		//I want to print characters side by side, so I must create the row loop to capture the ASCII height
		for row := 0; row < 8; row++ {
			//loop through each char of the input string (to get the rune) --> horizontal control
			for _, char := range val {
				start := (int(char) - 32) * 9 // calculate start for each character
				//lookup the banner file(splitfile) and print for each row of the char.
				result += lines[start+row]
			}
			result += "\n" // after finishing one row, add a new line so it prints the next row on a new line
		}
	}

	fmt.Print(result)
}
