package _interface

import "fmt"

type Person interface {
	Speak() string
	Run() string
}

type Man struct {
}

type WoMan struct {
}

func (d Man) Speak() string {
	return "我是个男人"
}

func (d Man) Run() string {
	return "我是个男人"
}

func personAllAction(p Person) {
	content := p.Speak()
	fmt.Printf("%v", content)
}

func Invok() {
	person := Man{}
	personAllAction(person)
}
