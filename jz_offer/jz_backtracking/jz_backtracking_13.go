package jz_backtracking

import (
//
)

// 描述
// 地上有一个 rows 行和 cols 列的方格。坐标从 [0,0] 到 [rows-1,cols-1] 。一个机器人从坐标 [0,0] 的格子开始移动，每一次只能向左，右，上，下四个方向移动一格，但是不能进入行坐标和列坐标的数位之和大于 threshold 的格子。 例如，当 threshold 为 18 时，机器人能够进入方格   [35,37] ，因为 3+5+3+7 = 18。但是，它不能进入方格 [35,38] ，因为 3+5+3+8 = 19 。请问该机器人能够达到多少个格子？

// 数据范围： 0 \le threshold \le 15 \0≤threshold≤15  ，1 \le rows,cols \le 100 \1≤rows,cols≤100
// 进阶：空间复杂度 O(nm) \O(nm)  ，时间复杂度 O(nm) \O(nm)

func digitalSum(num int) int {
	var sum int
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func dfs2(threshold int, rows int, cols int, x int, y int, isVisited [][]bool) int {
	if x < 0 || y < 0 || x >= rows || y >= cols || isVisited[x][y] || digitalSum(x)+digitalSum(y) > threshold {
		return 0
	}
	isVisited[x][y] = true
	return 1 + dfs2(threshold, rows, cols, x-1, y, isVisited) + dfs2(threshold, rows, cols, x+1, y, isVisited) + dfs2(threshold, rows, cols, x, y-1, isVisited) + dfs2(threshold, rows, cols, x, y+1, isVisited)
}

func movingCount(threshold int, rows int, cols int) int {
	if threshold < 0 || rows <= 0 || cols <= 0 {
		return 0
	}
	isVisited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		isVisited[i] = make([]bool, cols)
	}
	return dfs2(threshold, rows, cols, 0, 0, isVisited)
}
