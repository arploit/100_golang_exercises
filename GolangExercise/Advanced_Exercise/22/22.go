// Write a program to compute the frequency of the words from the input.
//  The output should output after sorting the key alphanumerically.
//  Suppose the following input is supplied to the program:
// New to Python or choosing between Python 2 and Python 3? Read Python 2 or Python 3.
//  Then, the output should be: 2:2 3.:1 3?:1 New:1 Python:5 Read:1 and:1 between:1 choosing:1 or:2 to:1

package main

import "fmt"

func WrdFrequency(sentence string) map[string]int {
	obj := make(map[string]int)

	for _, char := range sentence {
		if obj[string((char))] != -1 {
			obj[string((char))] += 1
		} else {
			obj[string((char))] = 0
		}
	}
	return obj
}

func PrintFrequency(sentence string, freqMap map[string]int) {
	for _, char := range sentence {
		fmt.Println(string(char), freqMap[string(char)]) //, freqMap[rune(char)])
	}
}

func main() {
	x := "New to Python or choosing between Python 2 and Python 3? Read Python 2 or Python 3"
	freqMap := WrdFrequency(x)
	// fmt.Println(freqMap)
	PrintFrequency(x, freqMap)
}
