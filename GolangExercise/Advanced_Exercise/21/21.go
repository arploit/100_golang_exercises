package main

import (
	"fmt"
	"math"
)

type Movements struct {
	Direction string
	Step      int
}

const (
	LEFT  = "LEFT"
	RIGHT = "RIGHT"
	DOWN  = "DOWN"
	UP    = "UP"
)

func CalculateDistance(Movement Movements, currentPos []int) []int {
	switch Movement.Direction {
	case UP:
		currentPos[1] += Movement.Step
	case DOWN:
		currentPos[1] -= Movement.Step
	case LEFT:
		currentPos[0] -= Movement.Step
	case RIGHT:
		currentPos[0] += Movement.Step

	}

	return currentPos

}

func main() {
	var movements = []Movements{
		{LEFT, 5},
		{DOWN, 3},
		{LEFT, 3},
		{RIGHT, 2},
	}

	finalPosition := []int{
		0, 0,
	}

	for i := 0; i < len(movements); i++ {
		finalPosition = CalculateDistance(movements[i], finalPosition)
	}

	dis := math.Sqrt(float64(finalPosition[0]*finalPosition[0] + finalPosition[1]*finalPosition[1]))

	fmt.Println(dis)

}
