package dogo

import (
	"context"
)

type Controller struct {
	Ctx  *context.Context
}

type ControllerInterface interface {
	Init(ctx *context.Context)
	Index()
}

func (c *Controller)Init(ctx *context.Context) {
	c.Ctx = ctx
}

func (c *Controller)Index() {

}


