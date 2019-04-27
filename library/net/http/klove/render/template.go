package render

import (
	"net/http"
	"html/template"
	"fmt"
)

type Template struct {
	Name 	string
	Data 	interface{}
}

var (
	tmpList map[string]*template.Template = make(map[string]*template.Template)
)


func (t Template)SetHeader(w http.ResponseWriter, key string, value string) {
	SetHeader(w,key,value)
}
func (t Template)SetContentType(w http.ResponseWriter) {
	t.SetHeader(w,"Content-Type","text/html; charset=utf-8")
}

func (t Template)Render(w http.ResponseWriter) (err error) {
	t.SetContentType(w)
	return execTemplate(w,t)
}


//添加一个模板
func AddTemplate(name string,files ...string) (err error) {
	if _, ok := tmpList[name];ok {
		return fmt.Errorf("The template %s is exist", name)
	}
	tmpList[name] = template.Must(template.ParseFiles(files...))
	return
}
//获取所有模板
func GetTemplate() (interface{}) {
	return tmpList
}
//执行模板
func execTemplate(w http.ResponseWriter, t Template) (err error) {
	// Ensure the template exists in the map.
	tmpl, ok := tmpList[t.Name]
	if !ok {
		return fmt.Errorf("The template %s does not exist", t.Name)
	}

	err = tmpl.ExecuteTemplate(w, "base", t.Data)
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
	return
}