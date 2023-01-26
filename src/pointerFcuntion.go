package main

import "fmt"

func add2(n int) {
	n += 2
}
func add2pt(n *int) {
	*n += 2
}
func swap(a, b string) {
	a, b = b, a
}
func swap2ptr(a, b *string) {
	*a, *b = *b, *a
}
func swap2ptr2add(a, b *string) {
	a, b = b, a
}
func main() {
	n := 4
	add2(n) //子函数赋值不会影响富含函数的值 如果要在子函数里修改父级变量 需要通过地址修改
	fmt.Println(n)
	add2pt(&n) //传入的是父级变量的地址 在子函数中修改了父级变量的值
	fmt.Println(n)
	c, d := "k", "m" //通过短声明解构赋值两个变量
	swap(c, d)       //直接在子函数内交换真实值 由于没有返回 并不会生效
	fmt.Println(c, d)
	swap2ptr(&c, &d) //在子函数内交换变量指针 可以修改父级变量的值
	fmt.Println(c, d)
	swap2ptr2add(&c, &d) //在子函数内交换变量指针的地址 并不会修改指针的值
	fmt.Println(c, d)
}
