package main


import (
	// "os"
	"fmt"
	"./foo"
	"sync"
)

var wg sync.WaitGroup

var mutex sync.RWMutex
func main()  {
	
	// fmt.Printf("main SubData:%+v\n",&sub.SubData)
	// fmt.Printf("main SubBss:%+v\n",&sub.SubBss)
	// fmt.Printf("main Pid:%d\n",os.Getpid())

	// sub.Print()
	wg.Add(1)
	go func(){
		// mutex.Lock()
		foo.FooT.Plus()
		fmt.Printf("%d\n",foo.FooT.Get())
		wg.Done()
		// mutex.Unlock()
	}()
	wg.Add(1)
	go func(){
		foo.FooT.Sub()
		fmt.Printf("%d\n",foo.FooT.Get())
		wg.Done()
	}()

	wg.Wait()
}