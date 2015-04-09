package routers

import (
	"github.com/oikomi/PrivateCloudStorageServer/ServerWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
