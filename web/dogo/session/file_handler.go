package session

import (
	"os"
	"io/ioutil"
	// "github.com/ghoiufyia/WxApp.kindle/web/dogo/log"
	// "fmt"
	"time"
)

type FileHandler struct {
	lifetime int64
	savePath string
}
// Open(savePath string,sessionName string) (interface{})

func (s *FileHandler)Open(savePath string,sessionName string) (interface{}) {
	return true
}
// Close() (interface{})

func (s *FileHandler)Close() (interface{}) {
	return true
}

func (s *FileHandler)Read(sessionId string) (Store,error) {
	path := s.savePath+"/"+sessionId
	fileInfo,err := os.Stat(path)
	if fileInfo == nil {
		return nil,err
	}
	t := time.Now()
	if fileInfo.ModTime().Unix() > t.Unix() + s.lifetime {
		content,err := ioutil.ReadFile(path)
		if err != nil {
			return nil,err
		}
		return content
	}
	// fmt.Printf("%+v",fileInfo)
	// fmt.Printf("%+v",err)
	return ""
}

func (s *FileHandler)Write(sessionId string,data string) (error) {
	path := s.savePath+"/"+sessionId
	err := ioutil.WriteFile(path,[]byte(data),0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *FileHandler)Destroy(sessionId string) (error) {
	path := s.savePath+"/"+sessionId
	err := os.Remove(path)
	if err != nil {
		return nil
	}
	return nil
}
func (s *FileHandler)Gc() error {
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