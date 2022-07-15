package jz_other

import (
	"strings" 
)

// 描述
// 请实现一个函数，将一个字符串s中的每个空格替换成“%20”。
// 例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。

// 数据范围:0 \le len(s) \le 1000 \0≤len(s)≤1000 。保证字符串中的字符为大写英文字母、小写英文字母和空格中的一种。

func replaceSpace( s string ) string {
    return strings.Replace(s, " ", "%20", -1)
}