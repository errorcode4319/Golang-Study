package main

import "fmt"

const (
	RED uint64 = 1 << iota
	BLUE
	GREEN
)

func main() {
	const C int = 10

	D := C * 100
	//fmt.Println(&C)
	//Cannot Take the Address of Const
	fmt.Println(RED, BLUE, GREEN)
}
