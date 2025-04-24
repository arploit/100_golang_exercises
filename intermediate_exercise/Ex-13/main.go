package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculateWN(sentence string, number string) string {
	return sentence + `number`
}

func divideString(userInput string) (int, int) {
	fmt.Print(userInput, "\n")
	 numberofDigits, numberofChar := 0, 0
	for _, char := range userInput {
			if char < 65 && char > 47 {
			numberofDigits++
			continue
		}
		if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) {
			numberofChar++
		}

	}
	return numberofDigits, numberofChar
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	var userString,_ = rdr.ReadString('\n') 
	userString = strings.TrimSpace(userString) 
	TotalDigits, TotalChar := divideString(userString)
	fmt.Printf("Total Char: %v\n Total Digits: %v", TotalChar, TotalDigits)
}