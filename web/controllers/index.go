package controllers

import (
	// "net/http"
	// "encoding/json"
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
	"fmt"
	// "io"
)

type IndexController struct {
	dogo.Controller
	Name	string
	age		int32
}

func (i *IndexController)Index ()  {
	fmt.Printf("adsd===============================")
	// fmt.Println(i.Ctx)
	fmt.Printf("%v",i.Ctx.Request)
	// io.WriteString(i.Ctx.ResponseWriter, "Hello from a HandleFunc #2!\n")
	

}









// func (i *IndexController)Index (w http.ResponseWriter, r *http.Request)  {
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(200)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"msg": "ok",
// 	})
// }

// func (ih *IndexHandler) SetAge (age int32) (error) {
// 	ih.age = age
// 	return nil
// }

// func (ih *IndexHandler) GetAge () (int32,error) {
// 	return ih.age,nil
// }