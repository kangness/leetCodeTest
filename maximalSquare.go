package main


func maximalSquare(matrix [][]byte) int {
	if matrix == nil {
		return 0
	}
	xl := len(matrix[0])
	yl := len(matrix)
	var dp [][]int
	for i := 0; i < yl; i ++ {
		dp = append(dp,make([]int, xl))
	}
	if matrix[0][0] == 1 {
		dp[0][0] = 1
	}
	for i := 0; i < xl; i ++ {
		for j := 0; j < yl; j ++ {
			if matrix[i][j] != 1 {
				dp[i][j] = 0
				continue
			}
			if i > 0 && j > 0 {
				if matrix[i - 1][j] == 0 {
					matrix[i][j] = 1
					continue
				}
				if j + 1 < yl && matrix[i][j + 1] == 0 {
					matrix[i][j] = 1
				}
				if j + 1 < yl && matrix[i - 1][j + 1] == 0 {
					matrix[i][j] = 1
				}
			}

		}
	}
}
