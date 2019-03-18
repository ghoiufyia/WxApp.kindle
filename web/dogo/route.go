package dogo

import (
	"reflect"
	"net/http"
	"regexp"
	"strings"
	"encoding/json"
)
//单个路由
type route struct {
	Name		string
	Method		string
	Pattern		string
	Controller	reflect.Type
	Action		string
}
//一个路由组
type RouteGroup struct {
	Prefix		string
	Controller 	reflect.Type
}
//多个路由组
type RouteMap struct {
	Routes []route
}
//新生成路由组
func NewRouteMap() *RouteMap {
	return &RouteMap{
		Routes:make([]route, 0),
	}
}
//添加路由
func (rm *RouteMap)RegisterRouteGroup(name string,method string,pattern string,c ControllerInterface,action string) {
	controller := reflect.Indirect(reflect.ValueOf(c)).Type()
	r := route{}
	r.Name = name
	r.Method = method
	r.Pattern = pattern
	r.Controller = controller
	r.Action = action

	// rm.Routes = append(rm.Routes,rg)
	rm.Routes = append(rm.Routes,r)
}
func (rm *RouteMap)Router(name string,method string,pattern string,c ControllerInterface,action string) {
	rm.RegisterRouteGroup(name,method,pattern,c,action)
}

//实现http的server handler
func (rm *RouteMap)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Log.Info("%v",r.URL)
	var myRoute route
	// Log.Info("%+v",myRoute)

	requestPath := r.URL.Path
	// Log.Info("%s\n",requestPath)
	// 匹配静态文件，后可由nginx定向
	for url,path := range StaticDir {
		if strings.HasPrefix(requestPath,url) {
			file := path + requestPath[len(url):]
			http.ServeFile(w,r,file)
			return
		}
	}
	// 匹配路由
	for _,v := range rm.Routes {
		// Log.Info("%+v",v)
		matched,err := regexp.MatchString(requestPath, v.Pattern)
		if err == nil && matched {
			myRoute = v
			break
		}
	}

	Log.Info("%+v",myRoute)

	// 未匹配到路由
	if nil == myRoute.Controller {
		Log.Info("未找到请求")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"msg": "没有",
		})
		// fmt.Printf("未找到该路由")
		return
	}
	// 反射出controller新对象
	vc := reflect.New(myRoute.Controller)
	// 执行Init方法
	init := vc.MethodByName("Init")
	in := make([]reflect.Value, 1)
	ctx := &Context{
		ResponseWriter:w,
		Request:r,
	}
	in[0] = reflect.ValueOf(ctx)
	init.Call(in)

	// 执行方法
	in = make([]reflect.Value, 0)
	index := vc.MethodByName(myRoute.Action)
	index.Call(in)

	// 扫尾方法
	in = make([]reflect.Value, 0)
	finish := vc.MethodByName("Finish")
	finish.Call(in)
	return
}