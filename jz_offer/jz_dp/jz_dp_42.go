package jz_dp

import (
	"math"
)

// 描述
// 输入一个长度为n的整型数组array，数组中的一个或连续多个整数组成一个子数组，子数组最小长度为1。求所有子数组的和的最大值。
// 数据范围:
// 1 <= n <= 2\times10^51<=n<=2×10 5
//  -100 <= a[i] <= 100−100<=a[i]<=100

// 要求:时间复杂度为 O(n)O(n)，空间复杂度为 O(n)O(n)
// 进阶:时间复杂度为 O(n)O(n)，空间复杂度为 O(1)O(1)

func FindGreatestSumOfSubArray(array []int) int {
	var sum int
	max := array[0]
	for _, v := range array {
		sum = int(math.Max(float64(sum+v), float64(v)))
		max = int(math.Max(float64(sum), float64(max)))
	}
	return max
}
