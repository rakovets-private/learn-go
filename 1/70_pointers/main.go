package main

import "fmt"

func main() {
	i := 5

	var b *int
	b = &i
	*b = 3
	println(i)

	c := &i
	*c = 6
	println(i)

	sl := []int{0, 0, 0}
	fmt.Println(sl)
	Tt(sl)
	fmt.Println(sl)
}

func Tt(items []int) {
	items[2] = 1
}
