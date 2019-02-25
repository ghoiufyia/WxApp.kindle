package handlers

import (
	"net/http"
	"encoding/json"
)

type IndexHandler struct {
	Name	string
	age		int32
}

func Index (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"msg": "ok",
	})
}

func (ih *IndexHandler) SetAge (age int32) (error) {
	ih.age = age
	return nil
}

func (ih *IndexHandler) GetAge () (int32,error) {
	return ih.age,nil
}