package jz_other

import (
//
)

// 描述
// 输入一个长度为 n 整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前面部分，所有的偶数位于数组的后面部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。

// 数据范围：0 \le n \le 50000≤n≤5000，数组中每个数的值 0 \le val \le 100000≤val≤10000
// 要求：时间复杂度 O(n)O(n)，空间复杂度 O(n)O(n)
// 进阶：时间复杂度 O(n^2)O(n 2)，空间复杂度 O(1)O(1)

func reOrderArray(array []int) []int {
	var odd, even []int
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 1 {
			odd = append(odd, array[i])
		} else {
			even = append(even, array[i])
		}
	}
	return append(odd, even...)
}
