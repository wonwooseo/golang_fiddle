package main

import (
	"fmt"
	linkedlist "golang_fiddle/structure_fiddle"
)

func main() {
	l := new(linkedlist.LinkedList)
	fmt.Println(l.Length, l.Head)
	l.Push(1)
	fmt.Println(l.Length, l.Head)
	l.Push(100)
	fmt.Println(l.Length, l.Head)
	l.Pop()
	fmt.Println(l.Length, l.Head)
}
