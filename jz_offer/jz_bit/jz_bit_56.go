package jz_bit

import (
//
)

// 描述
// 一个整型数组里除了两个数字只出现一次，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。

// 数据范围：数组长度 2\le n \le 10002≤n≤1000，数组中每个数的大小 0 < val \le 10000000<val≤1000000
// 要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

func FindNumsAppearOnce(array []int) []int {
	tmp, mask := 0, 1
	for _, v := range array {
		tmp ^= v
	}
	for tmp&mask == 0 {
		mask <<= 1
	}
	var a, b, c int
	for _, v := range array {
		if v&mask == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	if a > b {
		c = a
		a = b
		b = c
	}
	return []int{a, b}
}
