package jz_search

import (
// "brokenUd/golang/jz_offer/proto"
// "strings"
)

// 描述
// 给定一个长度为 n 的数组 nums 和滑动窗口的大小 size ，找出所有滑动窗口里数值的最大值。

// 例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，
// 那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}；
// 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}，
// {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。

// 数据范围： 1 \le size \le n \le 100001≤size≤n≤10000，数组中每个元素的值满足 |val| \le 10000∣val∣≤10000
// 要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)

func maxInWindows(num []int, size int) []int {
	if len(num) == 0 || size < 1 || size > len(num) {
		return nil
	}
	var ans, tmp []int
	for i, n := range num {
		for len(tmp) != 0 && num[tmp[len(tmp)-1]] < n {
			tmp = tmp[:len(tmp)-1]
		}
		tmp = append(tmp, i)
		if tmp[0]+size <= i {
			tmp = tmp[1:]
		}
		if i+1 >= size {
			ans = append(ans, num[tmp[0]])
		}
	}
	return ans
}
