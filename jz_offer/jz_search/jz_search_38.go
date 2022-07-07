package jz_search

import (
// "brokenUd/golang/jz_offer/proto"
// "math"
)

// 描述
// 输入一个长度为 n 字符串，打印出该字符串中字符的所有排列，你可以以任意顺序返回这个字符串数组。
// 例如输入字符串ABC,则输出由字符A,B,C所能排列出来的所有字符串ABC,ACB,BAC,BCA,CBA和CAB。

// 数据范围：n < 10n<10
// 要求：空间复杂度 O(n!)O(n!)，时间复杂度 O(n!)O(n!)

func Permutation(str string) []string {
	n := len(str)
	if n <= 1 {
		return []string{str}
	}
	var ans []string
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		s1 := str[i : i+1]
		for _, s2 := range Permutation(str[:i] + str[i+1:]) {
			newStr := s1 + s2
			m[newStr]++
		}
	}
	for v := range m {
		ans = append(ans, v)
	}
	return ans
}
