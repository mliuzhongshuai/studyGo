package _struct

import (
	"fmt"
)

/**
人类结构体
*/
type Person struct {
	name   string
	age    int
	action func(s string) (a, b int)
}

/**
person 方法定义
*/
func (p *Person) talk() string {
	return fmt.Sprintf("我的名字是%v,我今年%v岁了!", p.name, p.age)
}

/**
测试方法
*/
func ShowPerson() {
	person := Person{"刘中帅", 30, func(s string) (a, b int) {
		return 1, 2
	}}

	fmt.Println(person.talk())
}
