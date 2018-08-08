package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"relation-graph/graphRelation/graphServer/models"
	"fmt"
	"relation-graph/graphRelation/graphServer/cache"
)

type FileController struct {
	beego.Controller
}



func (this *FileController) CreateFileLink() {
	cflRaw := models.CreateFileLinkRaw{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cflRaw); err == nil {
		cfl := cflRaw.ConvertRawDataToCreateFileLink()
		fmt.Println("CreateFileLink", cfl)
		userByte, _ := json.Marshal(cfl.User)
		fileByte, _ := json.Marshal(cfl.File)
		cflByte, _ := json.Marshal(cfl)
		err0 := cache.PublishMsg(userByte, cache.User)
		err1 := cache.PublishMsg(fileByte, cache.File)
		err2 := cache.PublishMsg(cflByte, cache.CreateFileLink)
		if err0 == nil && err1 == nil && err2 == nil {
			this.Data["json"] = models.Ok.Map()
		} else {
			this.Data["json"] = models.Error.Map()
		}
		this.ServeJSON()
	} else {
		this.Data["json"] = models.Error.Map()
		this.ServeJSON()
	}
}

func (this *FileController) ClickFileLink() {
	cfl := models.ClickFileLink{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cfl); err == nil {
		fmt.Println("ClickFileLink", cfl)
		userByte, _ := json.Marshal(cfl.User)
		fileByte, _ := json.Marshal(cfl.File)
		cflByte, _ := json.Marshal(cfl)
		err0 := cache.PublishMsg(userByte, cache.User)
		err1 := cache.PublishMsg(fileByte, cache.File)
		err2 := cache.PublishMsg(cflByte, cache.ClickFileLink)
		if err0 == nil && err1 == nil && err2 == nil {
			this.Data["json"] = models.Ok.Map()
		} else {
			this.Data["json"] = models.Error.Map()
		}
		this.ServeJSON()
	} else {
		this.Data["json"] = models.Error.Map()
		this.ServeJSON()
	}
}