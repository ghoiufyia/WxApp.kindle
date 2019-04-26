package render


import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	templates map[string]*template.Template
	// bufpool   *bpool.BufferPool
	loaded    = false
)

func (c *Controller)RenderJson() {
	responseWriter := c.Ctx.ResponseWriter
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(200)
	json.NewEncoder(responseWriter).Encode(c.Data)
}


func RenderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	loadTemplates()

	fmt.Printf("%+v\n",templates)


	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist", name)
	}

	// Create a buffer to temporarily write to and check if any errors were encounterd.
	// buf := bufpool.Get()
	// defer bufpool.Put(buf)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		return err
	}

	// // The X-Frame-Options HTTP response header can be used to indicate whether
	// // or not a browser should be allowed to render a page in a <frame>,
	// // <iframe> or <object> . Sites can use this to avoid clickjacking attacks,
	// // by ensuring that their content is not embedded into other sites.
	// w.Header().Set("X-Frame-Options", "deny")
	// // Set the header and write the buffer to the http.ResponseWriter
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// buf.WriteTo(w)
	return nil
}

func loadTemplates() {
	if loaded {
		return
	}
	templates = make(map[string]*template.Template)

	// bufpool = bpool.NewBufferPool(64)

	layoutTemplates := map[string][]string{
		"./views/layouts/base.html": {
			"./views/index/index.html",
			"./views/index/list.html",
		},
	}

	for layout, includes := range layoutTemplates {
		for _, include := range includes {
			files := []string{include, layout}
			templates[filepath.Base(include)] = template.Must(template.ParseFiles(files...))
		}
	}

	loaded = true


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

func (c *Controller)RenderString(str string) {
	responseWriter := c.Ctx.ResponseWriter
	responseWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")
	responseWriter.WriteHeader(200)
	responseWriter.Write([]byte(str))
}

func (c *Controller)RenderTemplate(tmpl string) {
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

	RenderTemplate(c.Ctx.ResponseWriter,tmpl,nil)

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
