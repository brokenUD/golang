package jz_dp

import (
	// 
)

// 描述
// 请实现一个函数用来匹配包括'.'和'*'的正则表达式。
// 1.模式中的字符'.'表示任意一个字符
// 2.模式中的字符'*'表示它前面的字符可以出现任意次（包含0次）。
// 在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但是与"aa.a"和"ab*a"均不匹配

// 数据范围:
// 1.str 只包含从 a-z 的小写字母。
// 2.pattern 只包含从 a-z 的小写字母以及字符 . 和 *，无连续的 '*'。
// 3. 0 \le str.length \le 26 \0≤str.length≤26 
// 4. 0 \le pattern.length \le 26 \0≤pattern.length≤26 

func match( str string ,  pattern string ) bool {
    if len(pattern) == 0 {
        return len(str) == 0
    }
    firstMatch := len(str) > 0 && (pattern[0] == str[0] || pattern[0] == '.')
    if len(pattern) == 1 {
        return len(str) == 1 && firstMatch
    }
    if pattern[1] == '*' {
        if !firstMatch {
            return match(str, pattern[2:])
        }
        return match(str, pattern[2:]) || match(str[1:], pattern)
    }
    return firstMatch && match(str[1:], pattern[1:])
}