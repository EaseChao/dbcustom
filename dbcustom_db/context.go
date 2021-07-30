package dbcustom_db

import (
	"github.com/EaseChao/dbcustom/dbcustom_strings"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

/**
	这里为了获取参数值，不同框架随意亲自搭建不从获取方式
	需要用上封装方法，可以自己重写当前go
 */

// you can change this context
type Context struct {
	ctx iris.Context
}

func NewContext(ctx iris.Context) Context {
	return Context{ctx: ctx}
}

// 获取GET参数
func (c Context) GetValue(name string) string {
	if dbcustom_strings.IsNotBlank(name){
		return c.ctx.FormValue(name)
	}
	return ""
}

// 获取POST/PUT/...参数
func (c Context) PostValue(pojo interface{}) error {
	return simple.ReadForm(c.ctx, pojo)
}