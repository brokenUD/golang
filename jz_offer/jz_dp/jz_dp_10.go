package jz_dp

import (
//
)

// 描述
// 大家都知道斐波那契数列，现在要求输入一个正整数 n ，请你输出斐波那契数列的第 n 项。
// 斐波那契数列是一个满足 fib(x)=\left\{ \begin{array}{rcl} 1 &
// {x=1,2}\\ fib(x-1)+fib(x-2) &{x>2}\\ \end{array} \right.fib(x)={ 1 fib(x−1)+fib(x−2)
// ​x=1,2x>2的数列
// 数据范围：1\leq n\leq 401≤n≤40
// 要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n) ，本题也有时间复杂度 O(logn)O(logn) 的解法

func Fibonacci(n int) int {
	a, b, c := 1, 1, 1
	for i := 2; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
