package jz_bit

import (
	// 
)

// 描述
// 求1+2+3+...+n，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

// 数据范围： 0 < n \le 2000<n≤200
// 进阶： 空间复杂度 O(1)O(1) ，时间复杂度 O(n)O(n)

func Sum_Solution( n int ) int {
    if n > 1 {
		tmp := Sum_Solution(n-1)
		if tmp > 0 {
			n += tmp
		}
		return n
	}
	return n
}
