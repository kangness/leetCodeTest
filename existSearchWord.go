package main

import (
	"fmt"
)

type pairs struct{ x, y int }

var dirs = []pairs{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := 0; i < h; i++ {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }()
		for _, d := range dirs {
			ni, nj := d.x+i, d.y+j
			if ni < 0 || ni >= h {
				continue
			}
			if nj < 0 || nj >= w {
				continue
			}
			if vis[ni][nj] {
				continue
			}

			if check(ni, nj, k+1) {
				return true
			}
		}
		return false
	}
	for i, row := range board {
		for j, _ := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	word := "ABCCED"
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//	board := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	//	board := [][]byte{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}
	fmt.Println(exist(board, word))
}
