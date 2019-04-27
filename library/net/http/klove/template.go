package klove

import (
	// "fmt"
	// "html/template"
	// "net/http"
	// "path/filepath"
)

var (
	loaded bool = true
	loaded1 bool = false
	Load bool = true
	Load1 bool = false
    Str string = "aaaa"
)

func SetLoaded (status bool) {
	loaded = status
	loaded1 = status
	Load = status
	Load1 = status
	Str = "sss"
}


// func loadTemplates() {
// 	if loaded {
// 		return
// 	}
// 	templates = make(map[string]*template.Template)

// 	// bufpool = bpool.NewBufferPool(64)

// 	layoutTemplates := map[string][]string{
// 		"./views/layouts/base.html": {
// 			"./views/index/index.html",
// 			"./views/index/list.html",
// 		},
// 	}

// 	for layout, includes := range layoutTemplates {
// 		for _, include := range includes {
// 			files := []string{include, layout}
// 			templates[filepath.Base(include)] = template.Must(template.ParseFiles(files...))
// 		}
// 	}

// 	loaded = true
// }
