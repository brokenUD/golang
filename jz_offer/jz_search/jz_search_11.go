package jz_search

import (
// "brokenUd/golang/jz_offer/proto"
// "math"
)

// 描述
// 有一个长度为 n 的非降序数组，比如[1,2,3,4,5]，将它进行旋转，
// 即把一个数组最开始的若干个元素搬到数组的末尾，变成一个旋转数组，
// 比如变成了[3,4,5,1,2]，或者[4,5,1,2,3]这样的。请问，给定这样一个旋转数组，求数组中的最小值。

// 数据范围：1 \le n \le 100001≤n≤10000，数组中任意元素的值: 0 \le val \le 100000≤val≤10000
// 要求：空间复杂度：O(1)O(1) ，时间复杂度：O(logn)O(logn)

func minNumberInRotateArray(rotateArray []int) int {
	if len(rotateArray) == 0 {
		return 0
	}
	i, j := 0, len(rotateArray)-1
	for i < j {
		m := (i + j) / 2
		if rotateArray[m] > rotateArray[j] {   // 反过来可以？
			i = m + 1
		} else if rotateArray[m] < rotateArray[i] {
			j = m
		} else {
			j -= 1
		}
	}
	return rotateArray[i]
}
