package main

import (
	"golang.org/x/tour/pic"
)

// Pic return slices of slices
func Pic(dx, dy int) [][]uint8 {
	r := [][]uint8{}

	for y := 0; y < dy; y++ {
		t := []uint8{} // 必须每次初始化 t
		for x := 0; x < dx; x++ {
			t = append(t, uint8((x+y)/2)) // 实现 (x+y)/2
		}
		r = append(r, t)
	}

	return r
}

func main() {
	// 输出字符串后
	// 用 `data:image/png;base64,` 替换 `IMAGE:`
	// 将其写入 HTML 中 img 标签 src 属性值
	// 如 pic.html ，浏览器打开查看图片
	pic.Show(Pic)
}
