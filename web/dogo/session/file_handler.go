package session

import (
	"os"
	"io/ioutil"
	// "github.com/ghoiufyia/WxApp.kindle/web/dogo/log"
	// "fmt"
	"time"
)

type FileSession struct {
	lifetime int64
	savePath string
}

func (s *FileSession)Open(savePath string,sessionName string) (interface{}) {
	return true
}
func (s *FileSession)Close() (interface{}) {
	return true
}
func (s *FileSession)Read(sessionId string) interface{} {
	path := s.savePath+"/"+sessionId
	fileInfo,_ := os.Stat(path)
	if fileInfo == nil {
		return ""
	}
	t := time.Now()
	if fileInfo.ModTime().Unix() > t.Unix() + s.lifetime {
		content,err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		return content
	}
	// fmt.Printf("%+v",fileInfo)
	// fmt.Printf("%+v",err)
	return ""
}
func (s *FileSession)Write(sessionId string,data string) (error) {
	path := s.savePath+"/"+sessionId
	err := ioutil.WriteFile(path,[]byte(data),0644)
	if err != nil {
		return err
	}
	return nil
}
func (s *FileSession)Destroy(sessionId string) (error) {
	path := s.savePath+"/"+sessionId
	err := os.Remove(path)
	if err != nil {
		return nil
	}
	return nil
}
func (s *FileSession)Gc() error {
	f,err := os.Open(s.savePath)
	if err != nil {
		return err
	}
	fileList,err := f.Readdir(0)
	if err != nil {
		return err
	}
	t := time.Now()
	for _,fl := range fileList {
		if fl.ModTime().Unix() > t.Unix() + s.lifetime {
			os.Remove(s.savePath+fl.Name())
		}
	}
	return nil
}