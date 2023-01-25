package main

import "fmt"

func main() {
	index := 1
	for {
		fmt.Println("loopfow")
		break
	}
	for j := 7; j < 9; j++ {
		fmt.Println(j) //7-(9-1)
	}
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	for index <= 3 {
		fmt.Println(index)
		index += 1
	}
	for k := 0; k < 5; k++ {
		for j := 0; j < 5; j++ {
			fmt.Println("k = %d,j = %d\n", k, j)
		}
	}
}
