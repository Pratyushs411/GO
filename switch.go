package main

import (
	"fmt"
)

type s1 struct {
	name string
	age  int
}
type s2 struct {
	a int
	b int
}
type inter interface {
	check()
}

func (s11 s1) check() {
	fmt.Println("name is %s", s11.name)
}
func (s21 s2) check() int {
	return s21.a + s21.b
}

func swi(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Println("it is an int %t \n", v)
	default:
		fmt.Println("this is default \n")
	}
}
