package sessionhandler

import (
	"os"
	"io/ioutil"
	"time"
)
//文件的session handler
type FileHandler struct {
	Lifetime int64
	SavePath string
}
//打开文件
func (s *FileHandler)Open(savePath string,sessionName string) (interface{}) {
	return true
}
//关闭文件
func (s *FileHandler)Close() (interface{}) {
	return true
}
//读取session
func (s *FileHandler)Read(sessionId string) (string,error) {
	path := s.SavePath+"/"+sessionId
	fileInfo,_ := os.Stat(path)
	if fileInfo == nil {
		return "",nil
	}
	content,err := ioutil.ReadFile(path)
	if err != nil {
		return "",err
	}
	return string(content),nil
	// t := time.Now()
	// if fileInfo.ModTime().Unix() > t.Unix() + s.lifetime {
	// 	content,err := ioutil.ReadFile(path)
	// 	if err != nil {
	// 		return "",err
	// 	}
	// 	return string(content),nil
	// }
	// fmt.Printf("%+v",fileInfo)
	// fmt.Printf("%+v",err)
	// return "",nil
}
//写session
func (s *FileHandler)Write(sessionId string,data string) (err error) {
	path := s.SavePath+"/"+sessionId
	err = ioutil.WriteFile(path,[]byte(data),0644)
	if err != nil {
		return 
	}
	return
}
//销毁session
func (s *FileHandler)Destroy(sessionId string) (error) {
	path := s.SavePath+"/"+sessionId
	err := os.Remove(path)
	if err != nil {
		return nil
	}
	return nil
}
//session回收
func (s *FileHandler)Gc() error {
	f,err := os.Open(s.SavePath)
	if err != nil {
		return err
	}
	fileList,err := f.Readdir(0)
	if err != nil {
		return err
	}
	t := time.Now()
	for _,fl := range fileList {
		if fl.ModTime().Unix() > t.Unix() + s.Lifetime {
			os.Remove(s.SavePath + fl.Name())
		}
	}
	return nil
}