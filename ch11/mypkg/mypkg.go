package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParams()
	MethodWithParams(float64)
	MethodWithRreturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParams() {
	fmt.Println("MethodWithoutParams called")
}

func (m MyType) MethodWithParams(f float64) {
	fmt.Println("MethodWithParams called with", f)
}

func (m MyType) MethodWithRreturnValue() string {
	return "hi from MethodWithRreturnValue"
}

func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
