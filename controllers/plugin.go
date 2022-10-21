package controllers

type PluginController struct {
	BaseController
}

// @router /market [get]
func (c *PluginController) GetMarketList() {
	c.Success()
}

// @router /subscribe [get]
func (c *PluginController) GetSubscribeList() {
	c.Success()
}

// @router /publish [get]
func (c *PluginController) GetPublishList() {
	c.Success()
}

// @router /script [get]
func (c *PluginController) GetScriptList() {
	c.Success()
}

// @router /script/:id [get]
func (c *PluginController) GetScriptItem() {
	c.Success()
}
