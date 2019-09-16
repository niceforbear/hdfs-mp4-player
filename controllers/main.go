package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}
