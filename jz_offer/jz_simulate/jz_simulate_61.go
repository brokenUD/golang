package jz_simulate

import (
	"math"
)

// 描述
// 现在有2副扑克牌，从扑克牌中随机五张扑克牌，我们需要来判断一下是不是顺子。
// 有如下规则：
// 1. A为1，J为11，Q为12，K为13，A不能视为14
// 2. 大、小王为 0，0可以看作任意牌
// 3. 如果给出的五张牌能组成顺子（即这五张牌是连续的）就输出true，否则就输出false。
// 4.数据保证每组5个数字，每组最多含有4个零，数组的数取值为 [0, 13]

// 要求：空间复杂度 O(1)O(1)，时间复杂度 O(nlogn)O(nlogn)，本题也有时间复杂度 O(n)O(n) 的解法

func IsContinuous(numbers []int) bool {
	flag, min_, max_ := 0, 14, 0
	for _, v := range numbers {
		if v == 0 {
			continue
		}
		min_ = int(math.Min(float64(min_), float64(v)))
		max_ = int(math.Max(float64(max_), float64(v)))
		if (flag & (1 << v)) != 0 {
			return false
		}
		flag |= (1 << v)
	}
	return max_-min_ < 5
}
