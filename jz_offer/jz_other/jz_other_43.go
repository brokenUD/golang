package jz_other

import (
	// 
)

// 描述
// 输入一个整数 n ，求 1～n 这 n 个整数的十进制表示中 1 出现的次数
// 例如， 1~13 中包含 1 的数字有 1 、 10 、 11 、 12 、 13 因此共出现 6 次

// 注意：11 这种情况算两次
// 数据范围： 1 \le n \le 30000 \1≤n≤30000 
// 进阶：空间复杂度 O(1) \O(1)  ，时间复杂度 O(lognn) \O(lognn) 

func NumberOf1Between1AndN_Solution( n int ) int {
    // write code here
    bitNum, count := 1, 0
    high, cur, low := n/10, n%10, 0
    for high != 0 || cur != 0 {
        if cur < 1 {
            // case 1: cur == 0
            // cur=0时，高位需要减去一位用于低位进行计算
            count += high * bitNum
        } else if cur == 1 {
            // case 2: cur == 1
            // 相当于高位+低位计算结果，即(high * bitNum) + (low + 1)   
            count += high*bitNum + low + 1
        } else {
            // case3: cur > 1
            // 相对于cur=0的情况，就不需要高位减去一位来计算低位的结果数了
            // 相当于(high * bitNum) + (低位数结果数)
            count += (high + 1) * bitNum
        }
        // low、cur、high都像左偏移一个位
        // bitNum表示cur的数位
        low = low + cur*bitNum
        high, cur = high/10, high%10
        bitNum = bitNum * 10
    }
    return count

}