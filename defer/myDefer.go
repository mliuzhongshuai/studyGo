package _defer

import "fmt"

func TestPanicDefer() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(fmt.Sprintf("程序发生了错误:%v", err)) // 这里的err其实就是panic传入的内容，55
		}
	}()
	if 1 == 1 {
		panic("发生错误了")
	}
}
