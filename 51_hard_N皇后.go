package main

//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	var colums = make(map[int]bool)
	var rightDiagonals = make(map[int]bool)
	var leftDiagnoals = make(map[int]bool)
	var chessBoard = make([]int, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = -1
	}
	backtrack(chessBoard, n, 0, colums, rightDiagonals, leftDiagnoals)
	return solutions
}

func backtrack(chessBoard []int, n, row int, columns, rightDiagonals, leftDiagonals map[int]bool) {
	//终止条件
	if n == row {
		board := generateBoard(chessBoard, n)
		solutions = append(solutions, board)
	}

	// 循环每1列
	for i := 0; i < n; i++ {
		//过滤不符合条件的情况
		//此列有皇后了
		if columns[i] {
			continue
		}
		//此列的右斜线有皇后
		rightDiagonal := row - i
		if rightDiagonals[rightDiagonal] {
			continue
		}
		// 此列的左斜线有皇后
		leftDiagonal := row + i
		if leftDiagonals[leftDiagonal] {
			continue
		}
		// 走到这里说明可以放皇后, 设置各个参数
		chessBoard[row] = i // 在第row行的第i列放皇后
		columns[i] = true   // 第i列已经放了皇后
		rightDiagonals[rightDiagonal] = true
		leftDiagonals[leftDiagonal] = true
		backtrack(chessBoard, n, row+1, columns, rightDiagonals, leftDiagonals)
		// 重置参数
		chessBoard[row] = -1
		delete(columns, i)
		delete(rightDiagonals, rightDiagonal)
		delete(leftDiagonals, leftDiagonal)
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
