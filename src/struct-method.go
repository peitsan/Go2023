package main

import "fmt"

//func (u user) checkPassword(password string) bool {
//	return u.password == password
//}
//
//func (u user) resetPassword(password string) {
//	u.password = password
//}

type user struct {
	name     string
	password string
}

func (u *user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func (u *user) printlePasswd() {
	fmt.Println(u.name, u.password)
}

func (b *user) swapPasswd(a *user) { //可以实现对指针的方法操作  通过指针传递交换值
	fmt.Println(a.name, b.name)
	a, b = b, a
	fmt.Println(a.name, b.name)
}

func main() {
	a := user{name: "wang", password: "1024"}
	b := user{name: "li", password: "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048")) // true
	a.printlePasswd()
	a.swapPasswd(&b)
}
