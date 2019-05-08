package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

//WordCount return a map of the counts of each "word"
func WordCount(s string) map[string]int {
	m := make(map[string]int) // 使用 make 初始化 map

	words := strings.Fields(s) // Fields 函数根据空白字符拆分字符串

	for _, word := range words {
		m[word]++ // word 不存在时， m[word] 值为 0
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
