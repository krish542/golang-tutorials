package main

import "fmt"

type direction int

const (
	north direction = iota + 1
	south
	east
	west
)

func (d direction) String() string {
	return [...]string{"North", "South", "East", "West"}[d-1]
}
func (d direction) EnumIndex() int {
	return int(d)
}
func enum_custom() {
	var dir = east
	fmt.Println(dir)
	fmt.Println(dir.EnumIndex())
	fmt.Println(dir.String())
}
