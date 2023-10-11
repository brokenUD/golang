package jz_bit

import (
	// "math"
)

// 描述
// 实现函数 double Power(double base, int exponent)，求base的exponent次方。

// 注意：
// 1.保证base和exponent不同时为0。
// 2.不得使用库函数，同时不需要考虑大数问题
// 3.有特殊判题，不用考虑小数点后面0的位数。

// 数据范围： |base| \le 100 \∣base∣≤100  ， |exponent| \le 100 \∣exponent∣≤100  ,保证最终结果一定满足 |val| \le 10^4 \∣val∣≤10 4
// 进阶：空间复杂度 O(1)\O(1)  ，时间复杂度 O(n)\O(n)

func Power( base float64 ,  exponent int ) float64 {
	if exponent < 0 {
		base = 1/base
		exponent = -exponent
	}
	// res := 1.0
	return Pow(base, exponent)
}

func Pow( base float64 ,  exponent int ) float64 {
    res := float64(1)
	for exponent > 0 {
		if exponent & 1 > 0 {
			res *= base
		}
		base *= base
		exponent = exponent >> 1
	}
	return res
}
