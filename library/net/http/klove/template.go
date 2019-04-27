package klove

import (
	// "fmt"
	// "html/template"
	// "net/http"
	// "path/filepath"
)

var (
	loaded bool = false
)

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