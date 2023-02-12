package queenattack2

import (
	"fmt"
	"strings"
)

func QueenAttack() {
	result := SolveNQueens(4)
	fmt.Println(result)

	fmt.Println("==============")
	obtacles := [][]int32{[]int32{5, 5}, []int32{4, 2}, []int32{2, 3}}
	result2 := queensAttack(5, 5, 4, 3, obtacles)
	fmt.Println(result2)

}

func SolveNQueens(n int) [][]string {
	columns := make(map[int]struct{})
	pos_diagonal := make(map[int]struct{})
	neg_diagonal := make(map[int]struct{})

	board := make([][]string, n)
	for row := range board {
		board[row] = make([]string, n)
		for col := 0; col < n; col++ {
			board[row][col] = "."
		}
	}
	result := [][]string{}
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			temp := make([]string, n)
			for idx, rowValues := range board {
				temp[idx] = strings.Join(rowValues, "")
			}
			result = append(result, temp)
			return
		}
		for col := 0; col < n; col++ {
			if _, exists := columns[col]; exists {
				continue
			}
			if _, exists := pos_diagonal[row-col]; exists {
				continue
			}
			if _, exists := neg_diagonal[row+col]; exists {
				continue
			}
			columns[col] = struct{}{}
			pos_diagonal[row-col] = struct{}{}
			neg_diagonal[row+col] = struct{}{}
			board[row][col] = "Q"
			dfs(row + 1)
			delete(columns, col)
			delete(pos_diagonal, row-col)
			delete(neg_diagonal, row+col)
			board[row][col] = "."

		}
	}
	dfs(0)
	return result
}

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	// Write your code here

	// var createCaturBoard = make([][]int32, n)

	// for i := 1; i <= int(n); i++ {

	// 	// var ceate

	// 	// createCaturBoard = append(createCaturBoard)

	// 	fmt.Println("x")
	// 	// for j := clq
	// 	for j := i; j < int(n); j-- {
	// 		fmt.Println("  ")
	// 		fmt.Println("y")

	// 	}

	// }

	board := make([][]string, n)
	for row := range board {
		board[row] = make([]string, n)
		for col := 0; col < int(n); col++ {
			board[row][col] = "."
		}
	}

	fmt.Println(board)

	return 0

}

// [ x x x x x ]
// [ x x x x x ]
// [ x x x x x ]
// [ x x x x x ]
// [ x x x x x ]
