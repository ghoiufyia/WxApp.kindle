package modules

import (
	// "net/http"
	// "log"
	// "fmt"
	// "github.com/gorilla/mux"
	// "github.com/urfave/negroni"
	// "github.com/ghoiufyia/WxApp.kindle/web/modules"

)

func main() {
	port := "6767"
	log.Println("Starting HTTP service at " + port)
	fmt.Printf("Starting %v\n", appName)
	
	service := modules.NewService("app",16)

	r := mux.NewRouter()
	service.RegisterRoutes(r,"/v1")
	http.Handle("/", r)
	
    err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

    if err != nil {
        log.Println("An error occured starting HTTP listener at port " + port)
        log.Println("Error: " + err.Error())
    // }
	
}