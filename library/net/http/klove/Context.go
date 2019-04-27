package klove

import (
	"net/http"
	"encoding/json"
)
// 保存一次请求
type Context struct{
	Request        	*http.Request
	ResponseWriter 	http.ResponseWriter
	Session			*Session
}
// 初始化context
func (ctx *Context) Set(rw http.ResponseWriter, r *http.Request) {
	ctx.Request = r
	ctx.ResponseWriter = rw
}
// 设置http response code
func (c *Context)SetHttpStatus(code int) {
	c.ResponseWriter.WriteHeader(code)
}

// 未发现路由
func (c *Context)NotFound() {
	w := c.ResponseWriter
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 not found"))
}

// func (c *Context)Render(code int) {
// 	w := c.ResponseWriter
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(200)
// 	json.NewEncoder(w).Encode()
// }

func (c *Context)RenderText() {
	w := c.ResponseWriter
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 not found"))
}

// 返回json数据
func (c *Context)RenderJson(data interface{}) {
	c.SetHttpStatus(http.StatusOK)
	c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(c.ResponseWriter).Encode(data)
}

func (c *Context) RenderTemplate(template string,data interface{}) {
	
}

