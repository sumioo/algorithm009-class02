package week6

/*
//错误例子一则， 不能以到达目的格子的下一层作为终止条件，否则会重复计算。ps:画下递归树
func backtrace(matrix [][]int, row int, column int, count *int) {
	if row >= len(matrix) || column >= len(matrix[0]) {
		if (row-1 == len(matrix)-1 && column == len(matrix[0])-1) || (row == len(matrix)-1 && column-1 == len(matrix[0])-1) {
			*count++
			// fmt.Println(*count)
		}
		return
	}
	if matrix[row][column] == 1 {
		return
	}
	backtrace(matrix, row, column+1, count)
	backtrace(matrix, row+1, column, count)
}

*/

func backtrace(matrix [][]int, row int, column int, count *int) {
	if row == len(matrix)-1 && column == len(matrix[0])-1 {
		*count++
		return
	}
	if matrix[row][column] == 1 {
		return
	}
	if row < len(matrix)-1 && column < len(matrix[0])-1 {
		backtrace(matrix, row, column+1, count)
		backtrace(matrix, row+1, column, count)
	} else if row < len(matrix)-1 {
		backtrace(matrix, row+1, column, count)
	} else if column < len(matrix[0])-1 {
		backtrace(matrix, row, column+1, count)
	}
}

func backtraceB(matrix [][]int, row int, column int, count *int) {
	if row >= len(matrix) || column >= len(matrix[0]) {
		return
	}

	if row == len(matrix)-1 && column == len(matrix[0])-1 {
		*count++
		return
	}

	if matrix[row][column] == 1 {
		return
	}
	backtraceB(matrix, row, column+1, count)
	backtraceB(matrix, row+1, column, count)
}

func countPathsBacktrace(matrix [][]int) int {
	count := 0
	backtraceB(matrix, 0, 0, &count)
	return count
}

func topDown(matrix [][]int, row int, column int) int {
	if row >= len(matrix) || column >= len(matrix[0]) || matrix[row][column] == 1 {
		return 0
	}

	if row == len(matrix)-1 && column == len(matrix[0])-1 {
		return 1
	}
	return topDown(matrix, row+1, column) + topDown(matrix, row, column+1)

}

func topDownMemo(matrix [][]int, row int, column int, memo map[[2]int]int) int {
	if row >= len(matrix) || column >= len(matrix[0]) || matrix[row][column] == 1 {
		return 0
	}

	if row == len(matrix)-1 && column == len(matrix[0])-1 {
		return 1
	}

	key := [2]int{row, column}
	if v, ok := memo[key]; ok {
		return v
	}

	v := topDownMemo(matrix, row+1, column, memo) + topDownMemo(matrix, row, column+1, memo)
	memo[key] = v
	return v
}

func countPathsDPTopDown(matrix [][]int) int {
	// return topDownM(matrix, 0, 0)
	return topDownMemo(matrix, 0, 0, map[[2]int]int{})
}

func countPathsDPBottomUp(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rowCount, columnCount := len(matrix), len(matrix[0])

	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := range dp {
		dp[i][columnCount-1] = 1
	}
	for i := range dp[0] {
		dp[rowCount-1][i] = 1
	}

	for i := rowCount - 2; i >= 0; i-- {
		for j := columnCount - 2; j >= 0; j-- {
			if matrix[i][j] == 0 {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[0][0]
}
