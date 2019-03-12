package dogo

import (
	// "context"
	"fmt"
	// "net/http"
	"html/template"
	"path"
)

type Controller struct {
	Ctx  *Context
}

type ControllerInterface interface {
	Init(ctx *Context)
	Index()
	Render(0)
}

func (c *Controller)Init(ctx *Context) {
	c.Ctx = ctx
	
}

func (c *Controller)Index() {
	// ctx.Input.Data()
	// fmt.Printf("%v",c.Ctx)
}


func (c *Controller)Render() {
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

