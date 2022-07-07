package jz_dp

import (
//
)

// 描述
// 一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

// 数据范围：1 \leq n \leq 401≤n≤40
// 要求：时间复杂度：O(n)O(n) ，空间复杂度： O(1)O(1)

func jumpFloor(number int) int {
	a, b, c := 1, 1, 1
	for i := 2; i <= number; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
