package jz_backtracking

import (
//
)

// 描述
// 请设计一个函数，用来判断在一个n乘m的矩阵中是否存在一条包含某长度为len的字符串所有字符的路径。路径可以从矩阵中的任意一个格子开始，
// 每一步可以在矩阵中向左，向右，向上，向下移动一个格子。如果一条路径经过了矩阵中的某一个格子，则该路径不能再进入该格子。
// 例如 \begin{bmatrix} a & b & c &e \\ s & f & c & s \\ a & d & e& e\\ \end{bmatrix}\quad
// 矩阵中包含一条字符串"bcced"的路径，但是矩阵中不包含"abcb"路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，
// 路径不能再次进入该格子。
// 数据范围：0 \le n,m \le 20\0≤n,m≤20 ,1\le len \le 25\1≤len≤25

func hasPath(matrix [][]byte, word string) bool {
	if len(matrix[0]) == 0 && len(matrix) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if dfs(matrix, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(matrix [][]byte, word string, i int, j int, index int) bool {

	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) || matrix[i][j] != word[index] {
		return false
	}
	if index == len(word)-1 {
		return true
	}
	tmp := matrix[i][j]
	matrix[i][j] = '.'
	res := dfs(matrix, word, i+1, j, index+1) || dfs(matrix, word, i-1, j, index+1) || dfs(matrix, word, i, j-1, index+1) || dfs(matrix, word, i, j+1, index+1)
	matrix[i][j] = tmp
	return res
}
