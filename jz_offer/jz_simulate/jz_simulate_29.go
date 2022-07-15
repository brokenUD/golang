package jz_simulate

import (
//
)

// 描述
// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，例如，如果输入如下4 X 4矩阵：
// [[1,2,3,4],
// [5,6,7,8],
// [9,10,11,12],
// [13,14,15,16]]
// 则依次打印出数字
// [1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10]
// 数据范围:
// 0 <= matrix.length <= 100
// 0 <= matrix[i].length <= 100

func printMatrix(matrix [][]int) []int {
	var res []int
	if matrix == nil {
		return res
	}
	n := len(matrix)
	left, right, up, down := 0, len(matrix[0])-1, 0, n-1
	for left <= right && up <= down {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if up > down {
			break
		}
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
