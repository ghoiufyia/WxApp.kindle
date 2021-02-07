package render

import (
	"net/http"
	"encoding/json"
)

type Json struct {
	Data 	interface{}
}

func (j Json)SetHeader(w http.ResponseWriter, key string, value string) {
	SetHeader(w,key,value)
}
func (j Json)SetContentType(w http.ResponseWriter) {
	j.SetHeader(w,"Content-Type","application/json; charset=utf-8")
}

func (j Json)Render(w http.ResponseWriter) (err error) {
	j.SetContentType(w)
	return writeJson(w,j)
}

func writeJson(w http.ResponseWriter,j Json) (err error) {
	json.NewEncoder(w).Encode(j.Data)
	// err = errors.WithStack(err)
	return 
}