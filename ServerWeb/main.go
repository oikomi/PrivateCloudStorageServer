package main

import (
	_ "github.com/oikomi/PrivateCloudStorageServer/ServerWeb/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

