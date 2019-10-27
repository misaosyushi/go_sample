package main

import (
	"fmt"
	"math"
	"time"
)

var hoge, fuga, piyo string = "hoge", "fuga", "piyo"

func main() {

	message := "Hello"
	const HELLO string = "HELLO"
	var i int

	fmt.Println(message)
	fmt.Println(HELLO)
	fmt.Printf(hoge)
	fmt.Printf(fuga)
	fmt.Printf(piyo)
	fmt.Println(i)

	fmt.Println("The time is", time.Now())
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
	fmt.Println(math.Pi)
}