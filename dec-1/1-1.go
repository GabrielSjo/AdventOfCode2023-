package main

import (
	"fmt"
	"os"
	"unicode"
)

func isNumeric(b byte) bool {
	return unicode.IsDigit(rune(b))
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	// Init buffer to store one char at a time, var to store sum
	buffer := make([]byte, 1)
	sum := 0
	var firstNum, lastNum, combinedNumber int
	foundFirst := false

	for {
		_, err := file.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading file:", err)
			break
		}

		if isNumeric(buffer[0]) {
			if !foundFirst {
				firstNum = int(buffer[0] - '0')
				foundFirst = true
			}
			lastNum = int(buffer[0] - '0')
		}

		if buffer[0] == '\n' {
			combinedNumber = firstNum*10 + lastNum
			sum = sum + combinedNumber
			foundFirst = false
		}
	}

	fmt.Println("Sum of combined numbers:", sum)
}
