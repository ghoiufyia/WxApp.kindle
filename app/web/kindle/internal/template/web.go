package template

import (
	"github.com/ghoiufyia/kindle/library/net/http/klove/render"
)

func Init() {

	files := []string{"./views/index/index.html","./views/layouts/base.html"}
	render.AddTemplate("index",files...)

	files = []string{"./views/kindle/bind.html","./views/layouts/base.html"}
	render.AddTemplate("kbind_form",files...)

	// var tmp = make(map[string][]string)
	// tmp = map[string][]string{
	// 	"./views/layouts/base.html": {
	// 		"./views/index/index.html",
	// 		"./views/index/list.html",
	// 	},
	// }

	// for layout, includes := range tmp {
	// 	for _, include := range includes {
	// 		files := []string{include, layout}
	// 		templates[filepath.Base(include)] = template.Must(template.ParseFiles(files...))

	// 	}
	// }



}