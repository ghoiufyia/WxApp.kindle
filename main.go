package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"io"
	"sync"
)

var wg sync.WaitGroup
// func main(){
// 	var area = []string{"bj","tj","hb","sx","nmg","ln","jl","hlj","sh","js","zj","ah","fj","jx","sd","hn","hub","hun","gd","gx","hain","cq","sc","gz","yn","xz","sxs","gs","qh","nx","xj","tw","xg","am"}
// 	for _,a := range area {
// 		url := "http://www.shifansheng.net/rule/"+a+"_list.htaccess"
// 		url1 := "http://www.shifansheng.net/rule/"+a+".htaccess"
// 		wg.Add(1)
// 		go spider(a,url)
// 		wg.Add(1)
// 		go spider(a,url1)
// 	}

	
// 	wg.Wait()
// }

func spider(area string,url string) (err error) {
	content,err := getUrlData(url)
	// fmt.Println(content)
	if err != nil {
		return err
	}
	filename := "./file/"+area+".txt"
	err = writeFileString("./file/"+area+".txt",content,0644)
	if err != nil {
		return err
	}
	fmt.Println(filename+" get success")
	err = nil
	wg.Done()
	return err
}


func writeFileString(filename string,data string,perm os.FileMode) (err error) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.WriteString(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}


func getUrlData(url string) (content string,err error) {
	resp,err := http.Get(url)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	data,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	content = string(data)
	return content,nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}