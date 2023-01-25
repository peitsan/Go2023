package main

import "fmt"

func main() {
	//make是go初始化变量的关键字 类似c++ 的new go本身也有new关键字 当我们想要获取指向某个类型的指针时可以使用 new 关键字
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3

	// 使用append在尾部添加元素
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]

	c := make([]string, len(s))
	// 将s复制给c 声明长度已知string 数组
	copy(c, s)
	fmt.Println(c) // [a b c d e f]

	fmt.Println(s[2:5]) // [c d e] 索引和python完全一样 第a+1 到b的元素
	fmt.Println(s[:5])  // [a b c d e] 第一个到第5个
	fmt.Println(s[2:])  // [c d e f]

	good := []string{"g", "o", "o", "d"} //解包短声明 放在一个params里
	fmt.Println(good)                    // [g o o d]
}
