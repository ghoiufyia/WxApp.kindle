package dogo

import (
	"net/http"
)

type Context struct{
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

func (ctx *Context) Set(rw http.ResponseWriter, r *http.Request) {
	ctx.Request = r
	ctx.ResponseWriter = rw
}