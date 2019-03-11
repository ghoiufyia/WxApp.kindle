package dogo

import (
	// "context"
	"fmt"
	// "net/http"
	
)

type Controller struct {
	Ctx  *Context
}

type ControllerInterface interface {
	Init(ctx *Context)
	Index()
}

func (c *Controller)Init(ctx *Context) {
	c.Ctx = ctx
	
}

func (c *Controller)Index() {
	// ctx.Input.Data()
	fmt.Printf("%v",c.Ctx)
}


