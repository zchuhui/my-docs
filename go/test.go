package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ByAge 定义了一个自定义类型，用于按照年龄排序
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }              // 实现sort.Interface接口
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }    // 实现sort.Interface接口
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age } // 实现sort.Interface接口

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	sort.Sort(ByAge(people))
	// 输出: [{Charlie 20} {Alice 25} {Bob 30}]
	ByAge(people).Swap(0, 1)
	fmt.Println(people)
	fmt.Println(ByAge(people).Less(0, 1))
	fmt.Println(ByAge(people).Len())
}
