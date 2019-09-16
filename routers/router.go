package routers

import (
	"github.com/astaxie/beego"
	"github.com/niceforbear/hdfs-mp4-player/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/file_list", &controllers.FileListController{}, "GET:Get")
	beego.Router("/play", &controllers.PlayController{}, "GET:Get")
	beego.Router("/stream", &controllers.PlayController{}, "GET:StreamV2")

	beego.Router("/api/file_list", &controllers.FileListController{}, "GET:ApiGet")
}
