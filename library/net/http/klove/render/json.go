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
	j.SetContentType(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(j)
	// err = errors.WithStack(err)
	return nil
}