package main

import "fmt"

// 定义一个结构体
type Rectangle struct {
	width, height float64
}

// 定义一个方法：计算面积
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// 定义一个方法：修改宽度
func (r Rectangle) SetWidth(newWidth float64) {
	r.width = newWidth
}

func main() {

	rect := Rectangle{width: 10, height: 5}
	rect.SetWidth(2000)
	fmt.Println("矩形的宽度是:", rect.width, "矩形的面积是:", rect.Area())
}
