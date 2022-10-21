package controllers

type AdminController struct {
	BaseController
}

// @router /login [post]
func (c *AdminController) Login() {
	c.Success()
}

// @router /logout [post]
func (c *AdminController) Logout() {
	c.Success()
}
