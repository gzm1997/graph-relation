package controllers

import (
	"github.com/astaxie/beego"
	"relation-graph/graphRelation/graphServer/models"
	"encoding/json"
	"fmt"
	"relation-graph/graphRelation/graphServer/cache"
)

type GroupController struct {
	beego.Controller
}

func (this *GroupController) CreateGroupShareLink() {
	cgsl := models.CreateGroupShareLink{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cgsl); err == nil {
		fmt.Println("CreateGroupShareLink", cgsl)
		userByte, _ := json.Marshal(cgsl.User)
		groupByte, _ := json.Marshal(cgsl.Group)
		cgslByte, _ := json.Marshal(cgsl)
		err0 := cache.PublishMsg(userByte, cache.User)
		err1 := cache.PublishMsg(groupByte, cache.Group)
		err2 := cache.PublishMsg(cgslByte, cache.CreateGroupShareLink)
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

func (this *GroupController) ClickGroupShareLink() {
	cgsl := models.ClickGroupShareLink{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cgsl); err == nil {
		fmt.Println("ClickGroupShareLink", cgsl)
		userByte, _ := json.Marshal(cgsl.User)
		groupByte, _ := json.Marshal(cgsl.Group)
		cgslByte, _ := json.Marshal(cgsl)
		err0 := cache.PublishMsg(userByte, cache.User)
		err1 := cache.PublishMsg(groupByte, cache.Group)
		err2 := cache.PublishMsg(cgslByte, cache.ClickGroupShareLink)
		if err0 == nil && err1 == nil && err2 == nil {
			this.Data["json"] = models.Ok.Map()
		} else {
			//fmt.Println("err", err)
			this.Data["json"] = models.Error.Map()
		}
		this.ServeJSON()
	} else {
		//fmt.Println("err", err)
		this.Data["json"] = models.Error.Map()
		this.ServeJSON()
	}
}