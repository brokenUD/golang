package jz_dp

import (
	"math"
)

// 描述
// 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
// 数据范围:
// \ \text{s.length}\le 40000 s.length≤40000

func lengthOfLongestSubstring(s string) int {
	mp := make(map[string]int)
	res := 0
	dp := make([]int, len(s)+1)
	for i := 1; i < len(s)+1; i++ {
		dp[i] = 1
		if _, ok := mp[s[i-1:i]]; !ok {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = int(math.Min(float64(dp[i-1]+1), float64(i-mp[s[i-1:i]])))
		}
		mp[s[i-1:i]] = i
		res = int(math.Max(float64(res), float64(dp[i])))
	}
	return res
}
