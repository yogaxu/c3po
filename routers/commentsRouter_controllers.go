package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:PluginController"] = append(beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:PluginController"],
        beego.ControllerComments{
            Method: "GetMarketList",
            Router: "/market",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:PluginController"] = append(beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:PluginController"],
        beego.ControllerComments{
            Method: "GetPublishList",
            Router: "/publish",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:PluginController"] = append(beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:PluginController"],
        beego.ControllerComments{
            Method: "GetScriptList",
            Router: "/script",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:PluginController"] = append(beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:PluginController"],
        beego.ControllerComments{
            Method: "GetScriptItem",
            Router: "/script/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:PluginController"] = append(beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:PluginController"],
        beego.ControllerComments{
            Method: "GetSubscribeList",
            Router: "/subscribe",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:SystemController"] = append(beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetBasicInfo",
            Router: "/basic",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:SystemController"] = append(beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetStorageList",
            Router: "/storage",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:SystemController"] = append(beego.GlobalControllerRouter["github.com/yogaxu/c3po/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetStorageItem",
            Router: "/storage/:key",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
