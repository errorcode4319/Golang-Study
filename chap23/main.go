package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf(
			"Invalid Argument f:%g", f)
	}
	return math.Sqrt(f), nil
}

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "Password is Too Short"
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil
}

func devide(a, b int) {
	if b == 0 {
		panic("b cannot be 0")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	line, err := ReadFile("data.txt")
	if err != nil {
		fmt.Println("Read File Failed")
	} else {
		fmt.Println(line)
	}

	sqrt, err := Sqrt(2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Sqrt = %v\n", sqrt)
	}

	err = RegisterAccount("myID", "myPw")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("Register Success")
	}

	devide(9, 3)
	devide(9, 0)
}
