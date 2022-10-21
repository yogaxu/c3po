// @APIVersion 1.0.0
// @Title C3PO BOT API
// @Description Just the API document for this boring bot
// @Contact 375624193@qq.com
// @TermsOfServiceUrl http://yogaxu.top
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/yogaxu/c3po/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("",
		beego.NSNamespace("/system",
			beego.NSInclude(
				&controllers.SystemController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
