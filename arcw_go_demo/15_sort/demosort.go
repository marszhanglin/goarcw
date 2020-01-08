package main

import (
	"fmt"
	"sort"
)

// 定义一个数组，别名为ByLength
type ByLength []string

// ByLength 实现sort的Len，Less和Swap的三个接口，
// Len：长度  Swap：位置交换  Less：自定义排序
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 自定义：希望字符串长度升序排列
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
