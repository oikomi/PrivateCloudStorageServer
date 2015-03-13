package main

import (
	_ "github.com/oikomi/PrivateCloudStorageServer/ServerApi/docs"
	_ "github.com/oikomi/PrivateCloudStorageServer/ServerApi/routers"

	"github.com/astaxie/beego"
	"github.com/oikomi/PrivateCloudStorageServer/ServerApi/controllers"
)

func main() {
	beego.Router("api/v1/storage", &controllers.StorageController{})
	beego.Run()
}
