package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var logger = logs.GetLogger()

type BaseController struct {
	beego.Controller
}
