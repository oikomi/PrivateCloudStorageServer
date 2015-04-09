package main

import (
	_ "github.com/oikomi/PrivateCloudStorageServer/ServerApi/docs"
	_ "github.com/oikomi/PrivateCloudStorageServer/ServerApi/routers"

	"github.com/astaxie/beego"
	"github.com/oikomi/PrivateCloudStorageServer/ServerApi/controllers"
)

func main() {
	beego.Router("/index", &controllers.WebAccessController{})
	beego.Router("api/v1/storage", &controllers.StorageController{})
	beego.SetStaticPath("/open", "/home/data")
	beego.SetStaticPath("/static", "/mh/mygo/src/github.com/oikomi/PrivateCloudStorageServer/ServerApi/static")
	beego.Run()
}
