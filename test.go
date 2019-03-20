package main

import (
	"time"
	"fmt"
	"os"
)

func main(){
	for {
		time.Sleep(time.Millisecond*1000*10)
		file := "./test.log";
		f,err := os.OpenFile(file,os.O_RDWR|os.O_CREATE|os.O_APPEND,0766);
		if err!=nil {
			panic(err)
		}
		t := time.Now()
		str := fmt.Sprintf("%s\n",t);
		f.WriteString(str)
	}
}
