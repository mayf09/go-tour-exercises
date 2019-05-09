package main

import (
	"github.com/Go-zh/tour/reader"
)

// MyReader implement a Reader
type MyReader struct{}

// 这个练习的题意比较难理解
// 其实实现的 MyReader 仅仅是字面意义上的 Reader
// 具体的操作并不是读入，而是修改
func (r MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil // 返回读入的个数，无错误
}

func main() {
	reader.Validate(MyReader{})
}
