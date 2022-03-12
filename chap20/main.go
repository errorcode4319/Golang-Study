package main

import (
	"fmt"
	"gostudy/Interface/fedex"
	"gostudy/Interface/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

type Reader interface {
	Read(n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

type Attacker interface {
	Attack()
}

func main() {
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("Book", koreaPostSender)

	fedexSender := &fedex.FedexSender{}
	SendBook("Book", fedexSender)

	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{15})

	var att Attacker // Runtime Error
	att.Attack()
}
