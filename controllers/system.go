package controllers

type SystemController struct {
	BaseController
}

// @router /system/basic [get]
func (c *SystemController) GetBasic() {
	c.Success()
}
