package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	var rows, cols [9][9]int
	var sanbox [3][4][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			cols[j][index]++
			sanbox[i/3][j/3][index]++
			if rows[i][index] > 1 || cols[j][index] > 1 || sanbox[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(board))

}
