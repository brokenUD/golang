package jz_dp

import (
	"math"
)

// 描述
// 假设你有一个数组prices，长度为n，其中prices[i]是股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
// 1.你可以买入一次股票和卖出一次股票，并非每天都可以买入或卖出一次，总共只能买入和卖出一次，且买入必须在卖出的前面的某一天
// 2.如果不能获取到任何利润，请返回0
// 3.假设买入卖出均无手续费

// 数据范围： 0 \le n \le 10^5 , 0 \le val \le 10^40≤n≤10 5 ,0≤val≤10 4
// 要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

func maxProfit(prices []int) int {
	var max int
	min := math.MaxInt
	for _, v := range prices {
		max = int(math.Max(float64(v-min), float64(max)))
		min = int(math.Min(float64(v), float64(min)))
	}
	return max
}
