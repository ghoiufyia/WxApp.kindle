package main

import (


	"fmt"
)


func main() {
	array := [][]int32{{1,2},{1,4},{5,4},{5,6},{10,10},{10,2}}
	fmt.Println(len(array))

	recyle(array,5)

}


func recyle(a [][]int32,x int32) {
	var length int32 = 0;
	fmt.Println(len(a))
	for i:=len(a)-1;i>=0;i-- {
		a[i][0] = a[i][0] - a[0][0];
		a[i][1] = a[i][1] - a[0][1];
		fmt.Println(a[i])
		length = a[i][0] + a[i][1];
	}
	fmt.Println(length)
}