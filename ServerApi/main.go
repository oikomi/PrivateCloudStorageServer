package main

import (
	_ "github.com/oikomi/PrivateCloudStorageServer/ServerApi/docs"
	_ "github.com/oikomi/PrivateCloudStorageServer/ServerApi/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
