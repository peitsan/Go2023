package main

import "fmt"

func main() {
	nums := []int{2, 3, 4} //[2,3,4]
	sum := 0
	for i, num := range nums { //声明两个变量 (index,item) from nums
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
		}
	}
	fmt.Println(sum) // 9

	m := map[string]string{"a": "A", "b": "B", "c": "C"}
	for k, v := range m {
		fmt.Println(k, v) //(key,value) from map(stringkey:stringvalue)
	}
	for k := range m {
		fmt.Println("key", k) //打印key
	}
}
