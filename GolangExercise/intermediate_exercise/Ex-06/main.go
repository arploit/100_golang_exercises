package main

import (
	"fmt"
	"log"
	"math"
)

func CalculateEconomicQuantity(C int, H int, D int)float64 {

	number := int((2 * C * D)/H);
	 
	return float64(math.Sqrt(float64(number))) 
}

func main() {
	var D int
	fmt.Println("Enter The value of D:")
	_, err := fmt.Scan(&D)
	if err != nil{
		log.Panic("Error occured ", err)
	}

	fmt.Printf("The value of calculation is %f",CalculateEconomicQuantity(50,30,D))
}

// https://github.com/cblte/100-golang-exercises/blob/main/exercises_intermediate.adoc#exercise-006