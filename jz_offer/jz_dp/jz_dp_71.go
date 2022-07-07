package jz_dp

import (
	"math"
)

// 描述
// 一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶(n为正整数)总共有多少种跳法。

// 数据范围：1 \le n \le 201≤n≤20
// 进阶：空间复杂度 O(1)O(1) ， 时间复杂度 O(1)O(1)

func jumpFloorII(number int) int {
	if number <= 1 {
		return 1
	}
	return int(math.Pow(2, float64(number-1)))
}
