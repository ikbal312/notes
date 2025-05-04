package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	elementValue int
}

func (element *Element) String() string {
	return strconv.Itoa(element.elementValue)
}

type Stack struct {
	elements     []*Element
	elementCount int
}

func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount = stack.elementCount + 1
}
func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}
	length := len(stack.elements)
	var element *Element = stack.elements[length-1]
	if length > 1 {
		stack.elements = stack.elements[:length-1]
	} else {
		stack.elements = stack.elements[0:]
	}
	stack.elementCount = len(stack.elements)
	return element
}
func main() {
	var stack *Stack = &Stack{}
	stack.New()
	var element1 *Element = &Element{3}
	var element2 *Element = &Element{4}
	var element3 *Element = &Element{5}
	var element4 *Element = &Element{6}

	stack.Push(element1)
	stack.Push(element2)
	stack.Push(element3)
	stack.Push(element4)

	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop())
}
