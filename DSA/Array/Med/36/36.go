//https://leetcode.com/problems/valid-sudoku/description/?envType=problem-list-v2&envId=array

package main

import "fmt"

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		isIRan := false
		for j := 0; j < 9; j++ {
			rowMap := make(map[string]bool)
			columMap := make(map[string]bool)
			//for row check
			if !isIRan {
				for k := 0; k < 9; k++ {
					val := string(board[k][i])
					if rowMap[val] && val != "." {
						return false
					} else {
						rowMap[val] = true
					}
				}
				isIRan = true
			}

			//for column check
			for k := 0; k < 9; k++ {
				val := string(board[j][k])
				if columMap[val] && val != "." {
					return false
				} else {
					columMap[val] = true
				}
			}

			//for 3 X 3 Check
			boxStartRow := (i / 3) * 3
			boxStartCol := (j / 3) * 3
			box := make(map[string]bool)
			for h := boxStartRow; h <= boxStartRow+2; h++ {
				for l := boxStartCol; l <= boxStartCol+2; l++ {
					val := string(board[h][l])
					if box[val] && val != "." {
						return false
					} else {
						box[val] = true
					}
				}
			}

		}

	}

	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board3 := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}

	fmt.Println(isValidSudoku(board))
	fmt.Println(isValidSudoku(board2))
	fmt.Println(isValidSudoku(board3))

}
