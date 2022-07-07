package jz_queue_stack

import (
// "brokenUd/golang/jz_offer/proto"
// "math"
)

// 描述
// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。
// 假设压入栈的所有数字均不相等。例如序列1,2,3,4,5是某栈的压入顺序，
// 序列4,5,3,2,1是该压栈序列对应的一个弹出序列，但4,3,5,1,2就不可能是该压栈序列的弹出序列。

// 1. 0<=pushV.length == popV.length <=1000
// 2. -1000<=pushV[i]<=1000
// 3. pushV 的所有数字均不相同

func IsPopOrder(pushV []int, popV []int) bool {
	var n, j int
	for _, num := range pushV {
		pushV[n] = num
		for n >= 0 && pushV[n] == popV[j] {
			j++
			n--
		}
		n++
	}
	return n == 0
}
