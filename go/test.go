package main

import "fmt"

func add(a int, b int) (x, y int) {
	x = a * b
	return x, y
}

func main() {
	var arr [3]int
	arr[1] = 3

	colors := []string{"red", "green"}
	colors = append(colors, "blue")

	phonebook := map[string]string{
		"1": "123456",
		"2": "323232323",
	}

	fmt.Println(arr[2], colors[0:3], phonebook["1"])
	var c, d = add(1, 2)
	fmt.Println(c, d)
}
