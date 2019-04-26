package klove

import (
	// "fmt"
	// "encoding/json"
)
// 控制器结构
type Controller struct {
	Ctx  *Context
	Data map[string]interface{}
}
// 控制器接口
type ControllerInterface interface {
	Init(ctx *Context)
	Finish()
	SetData(string,interface{})
}
//初始化
func (c *Controller)Init(ctx *Context) {
	c.Ctx = ctx
	c.Data = make(map[string]interface{})
}
//扫尾工作
func (c *Controller) Finish() {
	// c.Ctx.Session.Save()
}
// 写入页面的数据
func (c *Controller) SetData(key string,value interface{}) {
	c.Data[key] = value
}