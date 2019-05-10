package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	// 中序遍历
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10) // 因为知道测试 tree 是 10 个元素
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < cap(ch1); i++ {
		v1, v2 := <-ch1, <-ch2
		// fmt.Printf("%d %d\n", v1, v2)
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("%v\n", Same(tree.New(1), tree.New(2)))
}
