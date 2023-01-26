package main

import "fmt"

type user struct {
	name     string
	password string
}

// 值传递
func checkPasswd(u user, pwd string) bool {
	u.name = "wang"
	return u.password == pwd
}

// 指针传递
func checkPassword2(u *user, password string) bool {
	u.name = "test"
	return u.password == password
}

func main() {
	a := user{name: "wang", password: "1024"}
	b := user{"wang", "1024"} //按序号解构
	c := user{name: "wang"}
	//d := user{"wang"}  //非法写法 参数必须完全写完
	var d = make([]user, 2) //使用make 构造一个2个长度的对象数组
	d[0].name = "lpc"
	d[1].password = "2096"
	var e user //声明一个空user
	e.password = "1024"
	e.name = "wang" //可以通过访问成员赋值
	fmt.Println(a, b, c, d, e)
	fmt.Println(c.password) //password没有赋值时
	fmt.Println(a.name, a.password)
	fmt.Println(checkPasswd(a, "1024"))
	fmt.Println(checkPassword2(&c, "3096")) //只传一个地址进去
}
