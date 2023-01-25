package main

import "fmt"

func main() {
	var a = "string"
	var b, c int = 10068, 2
	var d bool = true
	var e float32
	var f float64
	g := float32(f) //短声明
	h := a + "apple"
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println(h)
}
