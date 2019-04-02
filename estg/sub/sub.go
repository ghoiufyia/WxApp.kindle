package sub

import (
	"fmt"
	"os"
)

var SubData int = 10
var SubBss int

func Print()  {
	fmt.Printf("SubData:%+v\n",&SubData)
	fmt.Printf("SubBss:%+v\n",&SubBss)

	fmt.Printf("Sub Pid:%d\n",os.Getpid())

}