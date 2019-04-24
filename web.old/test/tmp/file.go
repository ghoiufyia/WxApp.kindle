package tmp

import (
	"fmt"
)

type file struct {
	num	int
}


func (f *file)Test()  {
	fmt.Printf("file")
}