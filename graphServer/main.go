package main

import (
	_ "relation-graph/graphRelation/graphServer/routers"

	"github.com/astaxie/beego"
	"relation-graph/graphRelation/graphServer/cache"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	forever := make(chan bool)
	go func() {
		cache.GetMsg()
	}()
	go func() {
		beego.Run()
	}()
	<- forever
}
