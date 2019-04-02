package foo

import (
	// "fmt"
)

var (
	FooT *Foo = &Foo{
		num:10,
	}
)

type Foo struct {
	num	int32
}

func (f *Foo)Plus() {
	f.num = f.num+1
}

func (f *Foo)Sub() {
	f.num = f.num-1
}

func (f *Foo)Get()(int32) {
	return f.num
}