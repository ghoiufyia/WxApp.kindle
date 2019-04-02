package dogo

import (
	"fmt"
	// "html/template"
	// "path"
	"encoding/json"
	
)

type Controller struct {
	Ctx  *Context
	Data map[string]interface{}
}

type ControllerInterface interface {
	Init(ctx *Context)
	Index()
	Render()
	RenderJson()
	RenderJsonByData(data map[string]interface{})
	// RenderTemplate(data map[string]interface{})
	Finish()
	SetData(string,interface{})
}

func (c *Controller)Init(ctx *Context) {
	c.Ctx = ctx
	c.Data = make(map[string]interface{})
}

func (c *Controller)Index() {
}
func (c *Controller)Finish() {}

// 写入页面的数据
func (c *Controller) SetData(key string,value interface{}) {
	c.Data[key] = value
}

func (c *Controller)Render() {
	// var ViewsPath string = "./views"
	// var filenames []string;

	// filenames = append(filenames,path.Join(ViewsPath,"/layouts/header.tmpl"))
	// filenames = append(filenames,path.Join(ViewsPath,"/index/index.tmpl"))
	// filenames = append(filenames,path.Join(ViewsPath,"/layouts/footer.tmpl"))
	// t,err := template.ParseFiles(filenames...)
	// if err != nil {
	// 	fmt.Printf("%v",err)
	// }

	// err = t.ExecuteTemplate(c.Ctx.ResponseWriter,"content",nil)

	RenderTemplate(c.Ctx.ResponseWriter,"index.html",nil)

}

func (c *Controller)RenderJson() {
	responseWriter := c.Ctx.ResponseWriter
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(200)
	json.NewEncoder(responseWriter).Encode(c.Data)
}

func (c *Controller)RenderJsonByData(data map[string]interface{}) {
	responseWriter := c.Ctx.ResponseWriter
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(200)
	json.NewEncoder(responseWriter).Encode(data)
}

// func (c *Controller)RenderTemplate(data map[string]interface{}) {
// 	var ViewsPath string = "./"
// 	var filenames []string;
// 	filenames = append(filenames,path.Join(ViewsPath,"views/layouts/footer.tmpl"))
// 	filenames = append(filenames,path.Join(ViewsPath,"views/layouts/header.tmpl"))
// 	t,err := template.ParseFiles(filenames...)
// 	if err != nil {
// 		fmt.Printf("%v",err)
// 	}
// 	err = t.ExecuteTemplate(c.Ctx.ResponseWriter,"footer",nil)
// 	err = t.ExecuteTemplate(c.Ctx.ResponseWriter,"header",nil)
// }

func (c *Controller)Error() {
	// c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	// c.Ctx.ResponseWriter.WriteHeader(200)
	// json.NewEncoder(c.Ctx.ResponseWriter).Encode(map[string]interface{}{
	// 	"msg": "ok",
	// })
	fmt.Printf("未找到该路由")


}