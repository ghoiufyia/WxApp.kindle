package main

import (
	"reflect"
	"fmt"
)

type T struct {}

type M interface {
	Hello() string
}

func main() {
    // name := "Hello"
    t := &T{}
	// reflect.ValueOf(t).MethodByName(name).Call(nil)
	
	fmt.Println(t)


	ttt := reflect.TypeOf(t)
	fmt.Println(ttt)

	var m M
	// m = t

	temp := reflect.TypeOf(m)
	// temp := reflect.Indirect(reflect.ValueOf(&T{})).Type()
	fmt.Println(temp)

	// t.Add(&T{})

}

func (t *T) Do() {
    fmt.Println("hello")
}

func (t *T) Hello() string {
	return "hello"
}

func (t *T)Add(c M) {
	temp := reflect.Indirect(reflect.ValueOf(c)).Type()
	fmt.Println(temp)
}