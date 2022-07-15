package jz_bit

import (
//
)

// 描述
// 写一个函数，求两个整数之和，要求在函数体内不得使用+、-、*、/四则运算符号。

// 数据范围：两个数都满足 -10 \le n \le 1000−10≤n≤1000
// 进阶：空间复杂度 O(1)O(1)，时间复杂度 O(1)O(1)

func Add(num1 int, num2 int) int {
	add, sum := num1, num2
	for add != 0 {
		tmp := sum ^ add
		add = (sum & add) << 1
		sum = tmp
	}
	return sum
}
