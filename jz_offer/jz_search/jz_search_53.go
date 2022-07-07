package jz_search

import (
// "brokenUd/golang/jz_offer/proto"
// "strings"
)

// 描述
// 给定一个长度为 n 的非降序数组和一个非负数整数 k ，要求统计 k 在数组中出现的次数

// 数据范围：0 \le n \le 1000 , 0 \le k \le 1000≤n≤1000,0≤k≤100，
// 数组中每个元素的值满足 0 \le val \le 1000≤val≤100
// 要求：空间复杂度 O(1)O(1)，时间复杂度 O(logn)O(logn)

func GetNumberOfK(data []int, k int) int {
	return bisearch(data, float32(k)+0.5) - bisearch(data, float32(k)-0.5)
}

func bisearch(data []int, k float32) int {
	left := 0
	right := len(data) - 1
	for left <= right {
		mid := (left + right) / 2
		if float32(data[mid]) < k {
			left = mid + 1
		} else if float32(data[mid]) > k {
			right = mid - 1
		}
	}
	return left
}
