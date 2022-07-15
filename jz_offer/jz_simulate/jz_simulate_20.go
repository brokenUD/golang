package jz_simulate

import (
//
)

// 描述
// 请实现一个函数用来判断字符串str是否表示数值（包括科学计数法的数字，小数和整数）。

// 科学计数法的数字(按顺序）可以分成以下几个部分:
// 1.若干空格
// 2.一个整数或者小数
// 3.（可选）一个 'e' 或 'E' ，后面跟着一个整数(可正可负)
// 4.若干空格

// 小数（按顺序）可以分成以下几个部分：
// 1.若干空格
// 2.（可选）一个符号字符（'+' 或 '-'）
// 3. 可能是以下描述格式之一:
// 3.1 至少一位数字，后面跟着一个点 '.'
// 3.2 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
// 3.3 一个点 '.' ，后面跟着至少一位数字
// 4.若干空格

// 整数（按顺序）可以分成以下几个部分：
// 1.若干空格
// 2.（可选）一个符号字符（'+' 或 '-')
// 3. 至少一位数字
// 4.若干空格

// 例如，字符串['+100','5e2','-123','3.1416','-1E-16']都表示数值。
// 但是['12e','1a3.14','1.2.3','+-5','12e+4.3']都不是数值。

// 提示:
// 1.1 <= str.length <= 25
// 2.str 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，空格 ' ' 或者点 '.' 。
// 3.如果怀疑用例是不是能表示为数值的，可以使用python的print(float(str))去查看
// 进阶：时间复杂度O(n)\O(n) ，空间复杂度O(n)\O(n)

func isNumeric(str string) bool {
	state := []map[byte]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4},
		{'d': 2, '.': 4},
		{'d': 2, '.': 3, 'e': 5, ' ': 8},
		{'d': 3, 'e': 5, ' ': 8},
		{'d': 3},
		{'s': 6, 'd': 7},
		{'d': 7},
		{'d': 7, ' ': 8},
		{' ': 8},
	}
	p := 0
	for k, _ := range str {
		c := str[k]
		var t byte
		if '0' <= c && c <= '9' {
			t = 'd'
		} else if c == '-' || c == '+' {
			t = 's'
		} else if c == 'e' || c == 'E' {
			t = 'e'
		} else if c == '.' || c == ' ' {
			t = c
		} else {
			t = '?'
		}
		if _, ok := state[p][t]; !ok {
			return false
		}
		p = state[p][t]
	}
	return p == 2 || p == 3 || p == 7 || p == 8
}