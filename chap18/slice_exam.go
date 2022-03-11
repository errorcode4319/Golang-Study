package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[:]

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100

	fmt.Println("After change second element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice = append(slice, 500)

	fmt.Println("After append 500")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice2 := append([]int{}, slice...)
	fmt.Println(slice)
	slice2[1] = 3
	fmt.Println(slice2)

}
