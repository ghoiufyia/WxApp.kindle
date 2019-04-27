package render

import (
	"net/http"
)

type Render interface {
	Render(http.ResponseWriter) error
	SetHeader(w http.ResponseWriter,key string,value string)
}

func SetHeader(w http.ResponseWriter, key string, value string) {
	w.Header().Set(key, value)
}