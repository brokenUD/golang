package jz_other

import (
//
)

// 描述
// 给定一个数组 A[0,1,...,n-1] ,请构建一个数组 B[0,1,...,n-1] ,其中 B 的元素 B[i]=A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]（除 A[i] 以外的全部元素的的乘积）。程序中不能使用除法。（注意：规定 B[0] = A[1] * A[2] * ... * A[n-1]，B[n-1] = A[0] * A[1] * ... * A[n-2]）
// 对于 A 长度为 1 的情况，B 无意义，故而无法构建，用例中不包括这种情况。

// 数据范围：1 \le n \le 10 \1≤n≤10  ，数组中元素满足 |val| \le 10 \∣val∣≤10

func multiply(A []int) []int {
	var B = make([]int, len(A))
	for i := range B {
		B[i] = 1
	}
	for i := 1; i < len(A); i++ {
		B[i] = B[i-1] * A[i-1]
	}
	tmp := 1
	for i := len(A) - 1; i >= 0; i-- {
		B[i] *= tmp
		tmp *= A[i]
	}
	return B
}
