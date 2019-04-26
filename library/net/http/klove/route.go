package klove

import (
	"reflect"
	"net/http"
	"errors"
	"log"
	// "regexp"
	// "strings"
	// "encoding/json"
)
//单个路由
type route struct {
	method		string
	pattern		string
	controller	reflect.Type
	action		string
	name		string
}
//一个路由组
type RouteGroup struct {
	prefix		string
	routes		map[string]route
	SessionmManager			*SessionManager
}

//新生成路由组
func NewRouteGroup() *RouteGroup {
	return &RouteGroup{
		prefix:	"",
		routes:	make(map[string]route, 0),
	}
}

//设置prefix
func (rg *RouteGroup)Prefix(prefix string) {
	rg.prefix = prefix
}

//设置sessionmanager
func (rg *RouteGroup)SetSessionmManager(sm *SessionManager){
	rg.SessionmManager = sm
}

//添加路由
func (rg *RouteGroup)Route(method string, pattern string,c ControllerInterface,action string,name string) error{
	controller := reflect.Indirect(reflect.ValueOf(c)).Type()
	if rg.prefix != "" {
		pattern = rg.prefix + pattern
	}
	r := route{}
	r.method = method
	r.pattern = pattern
	r.controller = controller
	r.action = action
	r.name = name

	if _,ok := rg.routes[pattern];ok {
		return errors.New("route exists")
	}
	rg.routes[pattern] = r
	return nil
}

//实现http的server handler
func (rg *RouteGroup)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		myRoute route
		ctx	*Context
	)
	ctx = &Context{
		ResponseWriter:w,
		Request:r,
	}

	// 初始化每个请求,包括session
	if rg.SessionmManager != nil {
		sess := rg.SessionmManager.SessionStart(ctx)
		ctx.Session = sess
	}
	// 匹配路由
	pattern := r.URL.Path
	
	myRoute,ok := rg.routes[pattern]
	if !ok {
		ctx.NotFound()
		return
	}
	log.Printf("%+v\n",myRoute)

	// Log.Info("%s\n",requestPath)
	// 匹配静态文件，后可由nginx定向
	// for url,path := range StaticDir {
	// 	if strings.HasPrefix(requestPath,url) {
	// 		file := path + requestPath[len(url):]
	// 		http.ServeFile(w,r,file)
	// 		return
	// 	}
	// }

	// 反射出controller新对象
	vc := reflect.New(myRoute.controller)
	// 执行Init方法,prev
	init := vc.MethodByName("Init")
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(ctx)
	init.Call(in)

	// 执行路由方法
	in = make([]reflect.Value, 0)
	action := vc.MethodByName(myRoute.action)
	action.Call(in)

	// 执行扫尾方法,next
	in = make([]reflect.Value, 0)
	finish := vc.MethodByName("Finish")
	finish.Call(in)

	//session回写
	err := rg.SessionmManager.SessionFlush(ctx)
	if err != nil {
		log.Printf("%+v\n",err)
	}
	return
}