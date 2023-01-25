package main

import (
	"fmt"
	"math"
)

func main() {
	const a = "string"
	const b = "append" + a
	const h = 666
	const c float64 = h / 200
	fmt.Println(a, b, c, math.Sin(c/10000))
}
