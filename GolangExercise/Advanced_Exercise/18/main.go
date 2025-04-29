//https://github.com/zhiwehu/Python-programming-exercises/blob/master/100%2B%20Python%20challenging%20programming%20exercises%20for%20Python%203.md#question-18

package main

import (
	"fmt"
	"regexp"
)

func ValidatePassword(pass string){
	passWordRegex := regexp.MustCompile(`[A-Za-z0-9$#@]`)
	if len(pass) > 6 ||  len(pass) < 12{
		for _,char := range pass {
			if !passWordRegex.MatchString(string(char)){
				fmt.Print("Invalid password ", pass)
				return
			}
		}
	}
	fmt.Println("Validated ", pass)
}

func main(){
	userPassword := "abc@ABsew_"
	ValidatePassword(userPassword)
}