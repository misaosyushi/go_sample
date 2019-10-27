package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"

	hoge := "hoge"
	fuga, piyo := true, "piyo"

	fmt.Println(i, j, c, python, java)
	fmt.Println(hoge, fuga, piyo)
}
