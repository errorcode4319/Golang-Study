package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	a := &account{100}

	withdrawFunc(a, 30)

	a.withdrawMethod(30)

	fmt.Printf("%d \n", a.balance)

	var b myInt = 10
	fmt.Println(b.add(30))
}
