package controllers

type SystemController struct {
	BaseController
}

// @router /basic [get]
func (c *SystemController) GetBasicInfo() {
	c.Success()
}

// @router /storage [get]
func (c *SystemController) GetStorageList() {
	c.Success()
}

// @router /storage/:key [get]
func (c *SystemController) GetStorageItem() {
	c.Success()
}
