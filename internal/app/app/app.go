package app

import "fmt"

type Application struct {
	Test Test
}

type Test struct {
}

func (t Test) Print() {
	fmt.Println("hello world")
}
