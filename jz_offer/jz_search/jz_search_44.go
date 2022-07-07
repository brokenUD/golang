package jz_search

import (
// "brokenUd/golang/jz_offer/proto"
	"strconv"
)

// 描述
// 数字以 0123456789101112131415... 的格式作为一个字符序列，
// 在这个序列中第 2 位（从下标 0 开始计算）是 2 ，第 10 位是 1 ，第 13 位是 1 ，以此类题，请你输出第 n 位对应的数字。

// 数据范围： 0 \le n \le 10^9 \0≤n≤10 9

func findNthDigit(n int) int {
	dig, bottom, top := 1, 0, 10
	for n > (top-bottom)*dig {
		n -= (top - bottom) * dig
		dig++
		bottom, top = top, top*10
	}
	num := n/dig + bottom
	r := n % bottom
	ans, _ := strconv.Atoi(string(num)[r:r+1])
	return ans
}
