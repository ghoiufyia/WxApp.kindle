package render

import (
	"net/http"
	"encoding/json"
)

type Json struct {
	
}

func (j Json)SetHeader(w http.ResponseWriter, key string, value string) {
	SetHeader(w,key,value)
}
func (j Json)SetContentType(w http.ResponseWriter) {
	j.SetHeader(w,"Content-Type","application/json; charset=utf-8")
}

func (j Json)Render(w http.ResponseWriter) (err error) {
	
	return writeJson(w,j)
}

func writeJson(w http.ResponseWriter,j Json) (err error) {
	j.SetContentType(w)
	json.NewEncoder(w).Encode(j)
	// err = errors.WithStack(err)
	return 
}