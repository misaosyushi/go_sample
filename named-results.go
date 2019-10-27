package main

import "fmt"

var c, python, java bool

func split(sum int) (x, y int)  {
	x = sum * 2
	y = sum + 1
	return
}

func main() {
	fmt.Println(split(16))

	var i int
	fmt.Println(i, c, python, java)
}
