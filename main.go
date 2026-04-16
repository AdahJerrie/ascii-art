package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	bannerfile := "standard"
	if len(os.Args) > 2 {
		bannerfile = os.Args[2]
	}

	banner := fmt.Sprintf("%s.txt", bannerfile)

	data, err := os.ReadFile(banner)
	if err != nil {
		log.Fatal(err)
	}

	input := os.Args[1]
	char := string(data)
	slicedata := strings.Split(char, "\n")
	sliceinput := strings.Split(input, "\\n")
	result := ""

	for _, line := range sliceinput {
		for row := 0; row < 8; row++ { //I want to print characters side by side, so I must create the row loop to capture the ASCII height
			for _, runes := range line {
				result += slicedata[((int(runes)-32)*9+2)+row] //lookup the banner file(slicedata) and print for each row of the char.
			}
			result += "\n" // after finishing one row, add a new line so it prints the next row on a new line
		}
	}
	fmt.Println(result)
}
