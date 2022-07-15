package jz_sort

import (
	"sort"
)

// 描述
// 给定一个长度为 n 的可能有重复值的数组，找出其中不去重的最小的 k 个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4(任意顺序皆可)。
// 数据范围：0\le k,n \le 100000≤k,n≤10000，数组中每个数的大小0 \le val \le 10000≤val≤1000
// 要求：空间复杂度 O(n)O(n) ，时间复杂度 O(nlogn)O(nlogn)

func GetLeastNumbers_Solution(input []int, k int) []int {
	if k == 0 || k > len(input) {
		return nil
	}
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	return input[:k]
}
