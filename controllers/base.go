package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

type ApiResult struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *BaseController) Success(data ...interface{}) {
	c.Data["json"] = &ApiResult{
		Success: true,
		Message: "请求成功",
		Data:    data,
	}
	_ = c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) Fail(message string, data ...interface{}) {
	c.Data["json"] = &ApiResult{
		Success: false,
		Message: message,
		Data:    data,
	}
	_ = c.ServeJSON()
	c.StopRun()
}
