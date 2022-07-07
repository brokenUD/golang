package jz_dp

import (
//
)

// 描述
// 有一种将字母编码成数字的方式：'a'->1, 'b->2', ... , 'z->26'。
// 我们把一个字符串编码成一串数字，再考虑逆向编译成字符串。
// 由于没有分隔符，数字编码成字母可能有多种编译结果，例如 11 既可以看做是两个 'a' 也可以看做是一个 'k' 。但 10 只可能是 'j' ，因为 0 不能编译成任何结果。
// 现在给一串数字，返回有多少种可能的译码结果

// 数据范围：字符串长度满足 0 < n \le 900<n≤90
// 进阶：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)

func solve(nums string) int {
	if len(nums) == 0 || nums[0] == '0' {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		if nums[i] != '0' {
			dp[i] = dp[i-1]
		}
		num := (nums[i-1]-'0')*10 + nums[i] - '0'
		if num >= 10 && num <= 26 {
			if i == 1 {
				dp[i] += 1
			} else {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[len(nums)-1]
}
