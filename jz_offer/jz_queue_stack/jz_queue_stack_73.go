package jz_queue_stack

import (
	// "brokenUd/golang/jz_offer/proto"
	"strings"
)

// 描述
// 牛客最近来了一个新员工Fish，每天早晨总是会拿着一本英文杂志，写些句子在本子上。
// 同事Cat对Fish写的内容颇感兴趣，有一天他向Fish借来翻看，但却读不懂它的意思。
// 例如，“nowcoder. a am I”。后来才意识到，这家伙原来把句子单词的顺序翻转了，
// 正确的句子应该是“I am a nowcoder.”。Cat对一一的翻转这些单词顺序可不在行，你能帮助他么？

// 数据范围：1 \le n \le 100 \1≤n≤100
// 进阶：空间复杂度 O(n) \O(n)  ，时间复杂度 O(n) \O(n)  ，保证没有只包含空格的字符串

func ReverseSentence(str string) string {
	// write code here
	ss := strings.Split(str, " ")
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[len(ss)-1-i] = ss[len(ss)-1-i], ss[i]
	}
	return strings.Join(ss, " ")
}