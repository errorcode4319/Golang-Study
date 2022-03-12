package main

import (
	"fmt"
	"os"
)

func sum(nums ...int) int {
	sum := 0
	fmt.Printf("nums Type: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func Print(args ...interface{}) string {
	for _, arg := range args {
		switch arg.(type) {
		case bool:
			val := arg.(bool)
			fmt.Println("Type bool => ", val)
		case float64:
			val := arg.(float64)
			fmt.Println("Type Float64 => ", val)
		case int:
			val := arg.(int)
			fmt.Println("Type Int => ", val)
		default:
			fmt.Println("Unsupported Type")
		}
	}
	return ""
}

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	Print(123, false, 3.14, "test")

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	//실행순서는 반대
	defer fmt.Println("Defer Run")
	defer f.Close()
	defer fmt.Println("File Closed")

	fmt.Fprintln(f, "Hello World")

	fn := getOperator("*")

	result := fn(3, 4)
	fmt.Println(result)

}
