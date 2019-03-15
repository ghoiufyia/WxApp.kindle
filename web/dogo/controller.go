package dogo

import (
	"fmt"
	"html/template"
	"path"
	
)

type Controller struct {
	Ctx  *Context
}

type ControllerInterface interface {
	Init(ctx *Context)
	Index()
	Render()
	renderJson()
	renderTemplate()
	Finish()
}

func (c *Controller)Init(ctx *Context) {
	c.Ctx = ctx
}

func (c *Controller)Index() {
}
func (c *Controller)Finish() {}

func (c *Controller)Render() {
	var ViewsPath string = "./views"
	var filenames []string;
	filenames = append(filenames,path.Join(ViewsPath,"/layouts/header.tmpl"))
	filenames = append(filenames,path.Join(ViewsPath,"/layouts/footer.tmpl"))
	t,err := template.ParseFiles(filenames...)
	if err != nil {
		fmt.Printf("%v",err)
	}
	err = t.ExecuteTemplate(c.Ctx.ResponseWriter,"header",nil)
	err = t.ExecuteTemplate(c.Ctx.ResponseWriter,"footer",nil)
}

func (c *Controller)renderJson() {
	var ViewsPath string = "./"
	var filenames []string;
	filenames = append(filenames,path.Join(ViewsPath,"views/layouts/header.tmpl"))
	filenames = append(filenames,path.Join(ViewsPath,"views/layouts/footer.tmpl"))
	t,err := template.ParseFiles(filenames...)
	if err != nil {
		fmt.Printf("%v",err)
	}
	err = t.ExecuteTemplate(c.Ctx.ResponseWriter,"header",nil)
	err = t.ExecuteTemplate(c.Ctx.ResponseWriter,"footer",nil)
}

func (c *Controller)renderTemplate() {
	var ViewsPath string = "./"
	var filenames []string;
	filenames = append(filenames,path.Join(ViewsPath,"views/layouts/header.tmpl"))
	filenames = append(filenames,path.Join(ViewsPath,"views/layouts/footer.tmpl"))
	t,err := template.ParseFiles(filenames...)
	if err != nil {
		fmt.Printf("%v",err)
	}
	err = t.ExecuteTemplate(c.Ctx.ResponseWriter,"header",nil)
	err = t.ExecuteTemplate(c.Ctx.ResponseWriter,"footer",nil)
}

func (c *Controller)Error() {
	// c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	// c.Ctx.ResponseWriter.WriteHeader(200)
	// json.NewEncoder(c.Ctx.ResponseWriter).Encode(map[string]interface{}{
	// 	"msg": "ok",
	// })
	fmt.Printf("未找到该路由")


}