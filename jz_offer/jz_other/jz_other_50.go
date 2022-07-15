package jz_other

import (
//
)

// 描述
// 在一个长为 字符串中找到第一个只出现一次的字符,并返回它的位置, 如果没有则返回 -1（需要区分大小写）.（从0开始计数）

// 数据范围：0 \le n \le 100000≤n≤10000，且字符串只有字母组成。
// 要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)

func FirstNotRepeatingChar(str string) int {
    var mp = make(map[byte]int, len(str))
	for i := range str {
		mp[str[i]]++
	}
	for i := range str {
		if mp[str[i]] == 1 {
			return i
		}
	}
	return -1
}
