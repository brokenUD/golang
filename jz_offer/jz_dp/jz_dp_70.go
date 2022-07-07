package jz_dp

import (
//
)

// 描述
// 我们可以用 2*1 的小矩形横着或者竖着去覆盖更大的矩形。请问用 n 个 2*1 的小矩形无重叠地覆盖一个 2*n 的大矩形，从同一个方向看总共有多少种不同的方法？

// 数据范围：0 \le n \le 38 \0≤n≤38
// 进阶：空间复杂度 O(1)\O(1)  ，时间复杂度 O(n)\O(n)

// 注意：约定 n == 0 时，输出 0
// 比如n=3时，2*3的矩形块有3种不同的覆盖方法(从同一个方向看)：

func rectCover(number int) int {
	if number <= 2 {
		return number
	}
	a, b, c := 1, 2, 0
	for i := 3; i <= number; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
