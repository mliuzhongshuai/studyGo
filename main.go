package main

import (
	"fmt"
	"studyGo/channel"
	_ "studyGo/defer"
	_defer "studyGo/defer"
	"studyGo/http"
	_ "studyGo/http"
	_ "studyGo/interface"
	_interface "studyGo/interface"
	"studyGo/json"
	_struct "studyGo/struct"
)

func init() {
	fmt.Println("执行初始化函数!")
}

func main() {

	fmt.Printf("\n%v\n", "channel练习.....")
	channel.TestChannel()
	fmt.Printf("%v\n", "接口练习.....")
	_interface.Invok()
	fmt.Printf("\n%v\n", "结构体练习.....")
	_struct.ShowPerson()
	fmt.Printf("\n%v\n", "json练习.....")
	json.TestEncodJson()
	fmt.Printf("\n%v\n", "panic练习.....")
	_defer.TestPanicDefer()
	fmt.Printf("\n%v\n", "启动http引擎.....")
	http.EndpointStarter()
}
