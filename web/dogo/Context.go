package dogo

import (
	"net/http"
	"github.com/ghoiufyia/WxApp.kindle/web/dogo/session"
)

type Context struct{
	Request        	*http.Request
	ResponseWriter 	http.ResponseWriter
	Session			*session.Store
}

func (ctx *Context) Set(rw http.ResponseWriter, r *http.Request) {
	ctx.Request = r
	ctx.ResponseWriter = rw
}