package jz_queue_stack

import (
// "brokenUd/golang/jz_offer/proto"
// "math"
)

// 描述
// 定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数，
// 输入操作时保证 pop、top 和 min 函数操作时，栈中一定有元素。

// 此栈包含的方法有：
// push(value):将value压入栈中
// pop():弹出栈顶元素
// top():获取栈顶元素
// min():获取栈中最小元素

// 数据范围：操作数量满足 0 \le n \le 300 \0≤n≤300  ，输入的元素满足 |val| \le 10000 \∣val∣≤10000
// 进阶：栈的各个操作的时间复杂度是 O(1)\O(1)  ，空间复杂度是 O(n)\O(n)

// var stack1 []int
// var stack2 []int

func Push2(node int) {
	stack1 = append(stack1, node)
	if len(stack2) == 0 {
		stack2 = append(stack2, node)
	} else {
		if node < Min2() {
			stack2 = append(stack2, node)
		} else {
			stack2 = append(stack2, stack2[len(stack2)-1])
		}
	}
}

func Pop2() int {
	res := stack1[len(stack1)-1]
	stack1 = stack1[:len(stack1)-1]
	stack2 = stack2[:len(stack2)-1]
	return res
}

func Top2() int {
	return stack1[len(stack1)-1]
}

func Min2() int {
	return stack2[len(stack2)-1]
}
